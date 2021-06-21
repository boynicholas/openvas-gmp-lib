package command_test

import (
	"log"
	"testing"

	gmp "github.com/boynicholas/openvas-gmp-lib"
	"github.com/boynicholas/openvas-gmp-lib/command"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	g, err := gmp.NewGmp(GetGmpConfig())
	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}

	uid, err := g.CreateTask(command.NewCreateTask(&command.CreateTask{
		Name:      "185.252.79.75",
		Comment:   "Nothing detected auto created",
		Alterable: false,
		UsageType: command.Scan,
		Config: &command.CreateTaskConfig{
			Id: "0cfd9ebc-d85c-42f2-be85-631af5c6e200",
		},
		Target: &command.CreateTaskTarget{
			Id: "070f82bf-d80a-4072-ae90-0a4e6d702c4e",
		},
		Scanner: &command.CreateTaskScanner{
			Id: "08b69003-5fc2-4037-a479-93b440211c73",
		},
	}))

	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}

	assert.NotNil(t, uid)
}
