package command_test

import (
	"log"
	"testing"

	gmp "github.com/boynicholas/openvas-gmp-lib"
	"github.com/boynicholas/openvas-gmp-lib/command"
)

func TestModifyTask(t *testing.T) {
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

	err = g.ModifyTask(command.NewModifyTask(&command.ModifyTask{
		TaskId: "ee3f71e9-0894-48d3-9a01-9fc7a2843568",
		Config: &command.ModifyTaskConfig{
			Id: "f2284cae-87d4-4589-bc00-e034fafbc121",
		},
	}))

	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}
}
