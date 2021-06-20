package command_test

import (
	"log"
	"testing"

	gmp "github.com/boynicholas/openvas-gmp-lib"
	"github.com/boynicholas/openvas-gmp-lib/command"
	"github.com/stretchr/testify/assert"
)

func TestGetTargets(t *testing.T) {
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

	resp, err := g.GetTargets(command.NewGetTargets(&command.GetTargets{
		Filter: "hosts=5.103.137.146",
	}))

	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}

	assert.NotNil(t, resp)
}
