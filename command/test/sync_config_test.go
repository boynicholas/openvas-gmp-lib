package command_test

import (
	"log"
	"testing"

	gmp "lyml.me/openvas-gmp-lib"
	"lyml.me/openvas-gmp-lib/command"
)

func TestSyncConfig(t *testing.T) {
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

	err = g.SyncConfig(command.NewSyncConfig("78eceaec-3385-11ea-b237-28d24461215b"))
	if err != nil {
		log.Fatalln(err)
		t.FailNow()
		return
	}
}
