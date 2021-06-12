package command_test

import (
	"os"
	"strconv"

	gmp "lyml.me/openvas-gmp-lib"
	"lyml.me/openvas-gmp-lib/command"
)

func GetGmpConfig() gmp.GmpConfig {
	gmpPort, _ := strconv.ParseUint(os.Getenv("GmpPort"), 10, 16)
	return gmp.GmpConfig{
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
