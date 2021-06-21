package command_test

import (
	"fmt"
	"log"
	"testing"

	gmp "github.com/boynicholas/openvas-gmp-lib"
	"github.com/boynicholas/openvas-gmp-lib/command"
)

func TestGetSettings(t *testing.T) {
	g, err := gmp.NewGmp(GetGmpConfig())
	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}

	resp, err := g.GetSettings(command.NewGetSettingsWithSingle("78eceaec-3385-11ea-b237-28d24461215b"))
	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}

	fmt.Println(resp)

}
