package command_test

import (
	"log"
	"testing"

	gmp "github.com/boynicholas/openvas-gmp-lib"
	"github.com/boynicholas/openvas-gmp-lib/command"
)

func TestModifyScanner(t *testing.T) {
	g, err := gmp.NewGmp(GetGmpConfig())
	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}

	err = g.ModifyScanner(command.NewModifyScanner(&command.ModifyScanner{
		ScannerId: "f04ecbd5-6423-49b0-a92f-a99d2617dc31",
		Type:      2,
		Name:      "Openvas Scanner 3",
		Host:      "172.31.11.32",
		Port:      33445,
	}))

	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}
}
