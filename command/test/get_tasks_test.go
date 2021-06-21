package command_test

import (
	"log"
	"testing"

	gmp "github.com/boynicholas/openvas-gmp-lib"
	"github.com/boynicholas/openvas-gmp-lib/command"
	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
	g, err := gmp.NewGmp(GetGmpConfig())
	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}

	resp, err := g.GetTasks(command.NewGetAllTasks())
	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}

	assert.NotNil(t, resp)
}
