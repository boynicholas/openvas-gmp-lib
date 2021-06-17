# Openvas GMP lib With Golang

This library wraps the Open VAS GMP protocol and is used to interact with the **Greenbone Vulnerability Manager**.

Currently under development, it implements some of the common features of GMP 21.4 and the latest 21.10.

## Installation and Usage

You need to first download the package to your project
```shell
go get github.com/boynicholas/openvas-gmp-lib
```


Then you can easily call it

```golang
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

	// todo something.

	defer gmp.Close()
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

```

## Support

For related questions, you can use Issue to feedback to me