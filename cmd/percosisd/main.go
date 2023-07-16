package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	percosis "github.com/percosis-labs/percosis/v16/app"
	"github.com/percosis-labs/percosis/v16/app/params"
	"github.com/percosis-labs/percosis/v16/cmd/percosisd/cmd"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, percosis.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
