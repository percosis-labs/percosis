package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	percosis "github.com/percosis-labs/percosis/app"
	"github.com/percosis-labs/percosis/app/params"
	"github.com/percosis-labs/percosis/cmd/percosisd/cmd"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, percosis.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
