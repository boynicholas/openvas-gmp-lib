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

	"github.com/boynicholas/openvas-gmp-lib/command"
)

type Client struct {
	dialer tls.Dialer
	conn   net.Conn
	ctx    context.Context
	mutex  sync.Mutex
	reader *xml.Decoder
}

type clientConfig struct {
	caPath      string
	cliCertPath string
	cliKeyPath  string
}

func newClient(ctx context.Context, cfg *clientConfig) (*Client, error) {
	cert, err := tls.LoadX509KeyPair(cfg.cliCertPath, cfg.cliKeyPath)
	if err != nil {
		return nil, err
	}

	caBytes, err := ioutil.ReadFile(cfg.caPath)
	if err != nil {
		return nil, err
	}

	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(caBytes)
	if !ok {
		return nil, command.ErrInvalidGvmCACert
	}

	tlsCfg := &tls.Config{
		RootCAs:            clientCertPool,
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	if ctx == nil {
		ctx = context.Background()
	}

	dialer := tls.Dialer{
		NetDialer: new(net.Dialer),
		Config:    tlsCfg,
	}

	return &Client{
		ctx:    ctx,
		dialer: dialer,
	}, nil
}

func (c *Client) Connect(addr string, port uint16) error {
	conn, err := c.dialer.DialContext(c.ctx, "tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return err
	}

	c.conn = conn
	c.reader = xml.NewDecoder(c.conn)
	return nil
}

// Send message
func (c *Client) Send(datas []byte, v interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, err := c.conn.Write(datas)
	if err != nil {
		return err
	}

	return c.read(v)
}

func (c *Client) read(val interface{}) error {
	token, err := c.reader.Token()
	if err != nil {
		return err
	}

	startElement, ok := token.(xml.StartElement)
	if !ok {
		return nil
	}

	if startElement.Name.Local == "gmp_response" {
		resp := &command.GmpResponse{}
		err = c.reader.DecodeElement(resp, &startElement)
		if err != nil {
			return err
		}

		if resp.Status != "200" {
			return errors.New(resp.StatusText)
		}

		// TODO: If it returns 200, I don’t know the actual operation
		return nil
	}

	err = c.reader.DecodeElement(val, &startElement)
	if err != nil {
		resp := &command.GmpResponse{}
		err = c.reader.Decode(resp)
		if err != nil {
			return err
		}
	}

	return nil
}
