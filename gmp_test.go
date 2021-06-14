package gmp

import (
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/boynicholas/openvas-gmp-lib/command"
)

func TestGmpNew(t *testing.T) {
	gmp, err := NewGmp(GetGmpConfig())

	if err != nil {
		log.Fatalln(err)
	}

	err = gmp.Authenticate(GetAuthenticate())
	if err != nil {
		log.Fatalln(err)
	}
}

func GetGmpConfig() GmpConfig {
	gmpPort, _ := strconv.ParseUint(os.Getenv("GmpPort"), 10, 16)
	return GmpConfig{
		Addr:              os.Getenv("GmpAddr"),
		Port:              uint16(gmpPort),
		TlsCaCertPath:     os.Getenv("GmpCaCert"),
		TlsClientCertPath: os.Getenv("GmpClientCert"),
		TlsClientKeyPath:  os.Getenv("GmpClientKey"),
	}
}

func GetAuthenticate() *command.Authenticate {
	return command.NewAuthenticate(os.Getenv("GmpUser"), os.Getenv("GmpPass"))
}
