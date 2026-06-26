// Command postgres-plugin runs the postgres engine as an out-of-process doze
// module — including its wire filter (TLS/startup/cancel) over the SCM_RIGHTS fd
// hand-off. The engine logic lives in this repo (modules/postgres/).
package main

import (
	"encoding/gob"

	"github.com/NerdMeNot/doze-modules/modules/postgres"
	dozeplugin "github.com/nerdmenot/doze-sdk/plugin"
)

func main() {
	gob.Register(&postgres.Config{})
	dozeplugin.Serve(postgres.Driver{})
}
