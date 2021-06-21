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

	err = g.ModifyTask(command.NewModifyTask(&command.ModifyTask{
		TaskId: "4aa27f6e-f1e9-45b0-b3b5-b552ab8c15fb",
		Config: &command.ModifyTaskConfig{
			Id: "b17204be-c413-47b0-bfc1-d1a776d4848d",
		},
		Scanner: &command.ModifyTaskScanner{
			Id: "f482bb74-41b2-4a1e-bd19-f4515f1e3b39",
		},
		Alert: []*command.ModifyTaskAlert{{
			Id: "1d3001d7-9ef5-432b-b47e-2785c706949a",
		}},
	}))

	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}
}
