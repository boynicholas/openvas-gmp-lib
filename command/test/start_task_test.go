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

	err = g.Authenticate(GetAuthenticate())
	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}

	uid, err := g.StartTask(command.NewStartTask("7733f39a-99e0-4c24-846e-4150bbd55357"))

	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}

	assert.NotNil(t, uid)
}
