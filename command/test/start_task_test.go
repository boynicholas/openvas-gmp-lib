package command_test

import (
	"log"
	"testing"

	gmp "github.com/boynicholas/openvas-gmp-lib"
	"github.com/boynicholas/openvas-gmp-lib/command"
	"github.com/stretchr/testify/assert"
)

func TestStartTask(t *testing.T) {
	g, err := gmp.NewGmp(GetGmpConfig())
	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}

	uid, err := g.StartTask(command.NewStartTask("ee3f71e9-0894-48d3-9a01-9fc7a2843568"))

	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}

	assert.NotNil(t, uid)
}
