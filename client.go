package gmp

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"sync"
	"time"

	"github.com/boynicholas/openvas-gmp-lib/command"
	"github.com/hashicorp/yamux"
)

type Client struct {
	dialer  tls.Dialer
	session *yamux.Session
	ctx     context.Context
	mutex   sync.Mutex
	cfg     *ClientConfig
}

type ClientConfig struct {
	CaPath       string
	CliCertPath  string
	CliKeyPath   string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	ConnTimeout  time.Duration
}

var DefaultTimeout time.Duration = time.Second * 10

func NewClient(ctx context.Context, cfg *ClientConfig) (*Client, error) {
	cert, err := tls.LoadX509KeyPair(cfg.CliCertPath, cfg.CliKeyPath)
	if err != nil {
		return nil, err
	}

	caBytes, err := ioutil.ReadFile(cfg.CaPath)
	if err != nil {
		return nil, err
	}

	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(caBytes)
	if !ok {
		return nil, command.ErrInvalidGvmCACert
	}

	tlsCfg := &tls.Config{
		RootCAs:      clientCertPool,
		Certificates: []tls.Certificate{cert},
	}

	if ctx == nil {
		ctx = context.Background()
	}

	dialer := tls.Dialer{
		NetDialer: new(net.Dialer),
		Config:    tlsCfg,
	}

	if cfg.ConnTimeout == 0 {
		cfg.ConnTimeout = DefaultTimeout
	}
	if cfg.WriteTimeout == 0 {
		cfg.WriteTimeout = DefaultTimeout
	}
	if cfg.ReadTimeout == 0 {
		cfg.ReadTimeout = DefaultTimeout
	}

	return &Client{
		cfg:    cfg,
		ctx:    ctx,
		dialer: dialer,
	}, nil
}

func (c *Client) Connect(addr string, port uint16) error {
	ctx, cancel := context.WithTimeout(c.ctx, c.cfg.ConnTimeout)
	defer cancel()
	conn, err := c.dialer.DialContext(ctx, "tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return err
	}
	session, err := yamux.Client(conn, nil)
	if err != nil {
		return err
	}
	c.session = session
	return nil
}

// Send message
func (c *Client) Send(datas []byte, v interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	stream, err := c.session.OpenStream()
	if err != nil {
		return err
	}
	defer stream.Close()
	err = stream.SetWriteDeadline(time.Now().Add(c.cfg.WriteTimeout))
	if err != nil {
		return err
	}

	_, err = stream.Write(datas)
	if err != nil {
		stream.Close()
		return err
	}

	// Create Reader
	err = stream.SetReadDeadline(time.Now().Add(c.cfg.ReadTimeout))
	if err != nil {
		return err
	}
	reader := xml.NewDecoder(stream)
	return c.read(reader, v)
}

func (c *Client) read(reader *xml.Decoder, val interface{}) error {
	token, err := reader.Token()
	if err != nil {
		return err
	}

	startElement, ok := token.(xml.StartElement)
	if !ok {
		return nil
	}

	if startElement.Name.Local == "gmp_response" {
		resp := &command.GmpResponse{}
		err = reader.DecodeElement(resp, &startElement)
		if err != nil {
			return err
		}

		if resp.Status != "200" {
			return errors.New(resp.StatusText)
		}

		// TODO: If it returns 200, I don’t know the actual operation
		return nil
	}

	err = reader.DecodeElement(val, &startElement)
	if err != nil {
		resp := &command.GmpResponse{}
		err = reader.Decode(resp)
		if err != nil {
			return err
		}
	}

	return nil
}
