package gmp

import (
	"log"
	"os"
	"strconv"
	"sync"
	"testing"

	"github.com/boynicholas/openvas-gmp-lib/command"
	"github.com/stretchr/testify/assert"
)

func TestGmpNew(t *testing.T) {
	gmp, err := NewGmp(GetGmpConfig())

	if err != nil {
		log.Fatalln(err)
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				resp, err := gmp.GetTargets(command.NewGetAllTargets())
				if !assert.NoError(t, err) {
					return
				}
				if !assert.NotNil(t, resp) {
					return
				}
				if !assert.NotNil(t, resp.Target) {
					return
				}
				if !assert.NotEmpty(t, resp.Target[0].Id) {
					return
				}
				t.Log("Get all targets request success\n")
			}

		}()
	}
	wg.Wait()
	// todo something.
}

func GetGmpConfig() GmpConfig {
	gmpPort, _ := strconv.ParseUint(os.Getenv("GmpPort"), 10, 16)
	return GmpConfig{
		Addr:              os.Getenv("GmpAddr"),
		Port:              uint16(gmpPort),
		TlsCaCertPath:     os.Getenv("GmpCaCert"),
		TlsClientCertPath: os.Getenv("GmpClientCert"),
		TlsClientKeyPath:  os.Getenv("GmpClientKey"),
		Username:          os.Getenv("GmpUser"),
		Password:          os.Getenv("GmpPass"),
	}
}

func GetAuthenticate() *command.Authenticate {
	return command.NewAuthenticate(os.Getenv("GmpUser"), os.Getenv("GmpPass"))
}
