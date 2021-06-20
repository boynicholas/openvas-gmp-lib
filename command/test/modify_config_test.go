package command_test

import (
	"log"
	"testing"

	gmp "github.com/boynicholas/openvas-gmp-lib"
	"github.com/boynicholas/openvas-gmp-lib/command"
)

func TestModifyConfig(t *testing.T) {
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

	err = g.ModifyConfig(command.NewModifyConfig(&command.ModifyConfig{
		ConfigId: "61aff01e-8f30-4011-a166-2756d040faeb",
		Scanner:  "f04ecbd5-6423-49b0-a92f-a99d2617dc31",
	}))

	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}
}
