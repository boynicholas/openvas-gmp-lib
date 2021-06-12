package command_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	gmp "lyml.me/openvas-gmp-lib"
	"lyml.me/openvas-gmp-lib/command"
)

func TestCreateTarget(t *testing.T) {
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

	target := command.NewCreateTarget(&command.CreateTarget{
		Name:      "test target",
		Comment:   "Test target comment",
		Hosts:     "127.0.0.1",
		PortRange: "80,443,3389,3306,22",
	})

	uid, err := g.CreateTarget(target)
	assert.NotNil(t, uid)
	assert.NoError(t, err)
}
