package command_test

import (
	"log"
	"testing"

	gmp "github.com/boynicholas/openvas-gmp-lib"
	"github.com/boynicholas/openvas-gmp-lib/command"
	"github.com/stretchr/testify/assert"
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
		Name:      "185.252.79.75",
		Comment:   "Nothing detected auto created",
		Hosts:     "185.252.79.75",
		PortRange: "21,22,80,143,443,465,587,993,2121,4333,8333,9900,18080,",
	})

	uid, err := g.CreateTarget(target)
	assert.NotNil(t, uid)
	assert.NoError(t, err)
}
