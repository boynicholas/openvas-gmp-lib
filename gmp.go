package gmp

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/boynicholas/openvas-gmp-lib/command"
	"github.com/google/uuid"
)

type Command interface {
	String() string
	Command() string
	GetRespStruct() interface{}
	Handler(v interface{}) (interface{}, error)
}

type GmpVersion string

const (
	// https://docs.greenbone.net/API/GMP/gmp-21.04.html
	GMP_21_04 GmpVersion = "21.4"
	// https://docs.greenbone.net/API/GMP/gmp-21.10.html
	GMP_21_10 GmpVersion = "21.10"
	// https://docs.greenbone.net/API/GMP/gmp-20.08.html
	GMP_20_08 GmpVersion = "20.8"

	Nil GmpVersion = ""
)

func GetGmpVersion(version string) GmpVersion {
	if version == string(GMP_21_04) {
		return GMP_21_04
	} else if version == string(GMP_20_08) {
		return GMP_20_08
	} else if version == string(GMP_21_10) {
		return GMP_21_10
	} else {
		return Nil
	}
}

type Gmp struct {
	version GmpVersion
	ctx     context.Context
	cfg     GmpConfig
	client  *Client
}

type GmpConfig struct {
	// Version of Greenbone Security Manager, default version is 21.04
	Version GmpVersion

	// Address of Greenbone Security Manager
	Addr string

	Port uint16

	// GVM CA certificate path
	TlsCaCertPath string
	// GVM Client certificate path
	TlsClientCertPath string
	// GVM Client key path
	TlsClientKeyPath string

	IsReconnect bool // Whether to reconnect when the connection is closed

	Username string // Cannot be empty when reconnect is enabled
	Password string // Cannot be empty when reconnect is enabled

	ctx context.Context
}

func NewGmp(cfg GmpConfig) (*Gmp, error) {
	if _, err := os.Stat(cfg.TlsCaCertPath); os.IsNotExist(err) {
		return nil, command.ErrGvmCACertNotFound
	}

	if _, err := os.Stat(cfg.TlsClientCertPath); os.IsNotExist(err) {
		return nil, command.ErrGvmClientCertNotFound
	}

	if _, err := os.Stat(cfg.TlsClientKeyPath); os.IsNotExist(err) {
		return nil, command.ErrGvmClientKeyNotFound
	}

	if cfg.ctx == nil {
		cfg.ctx = context.Background()
	}

	client, err := NewClient(cfg.ctx, &ClientConfig{
		CaPath:      cfg.TlsCaCertPath,
		CliCertPath: cfg.TlsClientCertPath,
		CliKeyPath:  cfg.TlsClientKeyPath,
	})

	if err != nil {
		return nil, err
	}

	if cfg.IsReconnect {
		if cfg.Username == "" || cfg.Password == "" {
			return nil, errors.New("when reconnection is enabled, the username and password cannot be empty")
		}
	}

	gmp := &Gmp{
		cfg:    cfg,
		client: client,
		ctx:    cfg.ctx,
	}

	err = gmp.init()
	if err != nil {
		return nil, err
	}

	return gmp, nil
}

func (g *Gmp) init() error {
	err := g.client.Connect(g.cfg.Addr, g.cfg.Port)
	if err != nil {
		return err
	}

	// Check version
	version, err := g.GetVersion()
	if err != nil {
		return err
	}

	v := GetGmpVersion(version)
	if v == Nil {
		return command.ErrUnsupportGVMVersion
	}

	g.version = v
	return nil
}

func (g *Gmp) Authenticate(authenticate *command.Authenticate) error {
	_, err := g.exec(authenticate, authenticate)
	if err != nil {
		return err
	}

	return nil
}

func (g *Gmp) GetVersion() (string, error) {
	v := command.NewGetVersion()
	r, err := g.exec(v, v)

	if err != nil {
		return "", err
	}

	return r.(string), nil
}

func (g *Gmp) CreateTarget(target *command.CreateTarget) (*uuid.UUID, error) {
	id, err := g.exec(target, target)
	if err != nil {
		return nil, err
	}

	uid, _ := uuid.Parse(id.(string))
	return &uid, nil
}

func (g *Gmp) CreateTask(task *command.CreateTask) (*uuid.UUID, error) {
	id, err := g.exec(task, task)

	if err != nil {
		return nil, err
	}

	uid, _ := uuid.Parse(id.(string))
	return &uid, nil
}

func (g *Gmp) GetSettings(filter *command.GetSettings) (*command.GetSettingsResp, error) {
	resp, err := g.exec(filter, filter)
	if err != nil {
		return nil, err
	}

	return resp.(*command.GetSettingsResp), nil
}

func (g *Gmp) StartTask(req *command.StartTask) (*uuid.UUID, error) {
	reportId, err := g.exec(req, req)
	if err != nil {
		return nil, err
	}

	uid, _ := uuid.Parse(reportId.(string))
	return &uid, nil
}

func (g *Gmp) SyncConfig(req *command.SyncConfig) error {
	_, err := g.exec(req, req)
	if err != nil {
		return err
	}

	return nil
}

func (g *Gmp) GetNvtFamilies(req *command.GetNvtFamilies) (*command.GetNvtFamiliesResp, error) {
	resp, err := g.exec(req, req)
	if err != nil {
		return nil, err
	}

	return resp.(*command.GetNvtFamiliesResp), nil
}

func (g *Gmp) exec(req interface{}, cmd Command) (interface{}, error) {
	data, err := xml.Marshal(req)
	if err != nil {
		return nil, err
	}

	rp := cmd.GetRespStruct()

	err = g.client.Send(data, rp)
	if err != nil {
		_, ok := err.(net.Error)
		if !ok {
			return nil, err
		}

		if g.cfg.IsReconnect {
			fmt.Printf("[Gmp] GMP connection error: %s, try to reconnect", err.Error())

			// reconnect
			err = g.reconnect()
			if err != nil {
				return nil, err
			} else {
				return g.exec(req, cmd)
			}
		} else {
			return nil, err
		}
	}

	return cmd.Handler(rp)
}

func (g *Gmp) reconnect() error {
	err := g.init()
	if err != nil {
		return err
	}

	err = g.Authenticate(command.NewAuthenticate(g.cfg.Username, g.cfg.Password))
	if err != nil {
		return err
	}

	return nil
}
