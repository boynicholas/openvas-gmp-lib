package command_test

import (
	"fmt"
	"log"
	"testing"

	gmp "github.com/boynicholas/openvas-gmp-lib"
	"github.com/boynicholas/openvas-gmp-lib/command"
)

func TestGetNvtFamilies(t *testing.T) {
	g, err := gmp.NewGmp(GetGmpConfig())
	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}

	err = g.Authenticate(GetAuthenticate())
	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}

	resp, err := g.GetNvtFamilies(command.NewGetNvtFamilies(""))
	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}

	fmt.Println(resp)

}
