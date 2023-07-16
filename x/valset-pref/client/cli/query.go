package valsetprefcli

import (
	"github.com/spf13/cobra"

	"github.com/percosis-labs/percosis/osmoutils/percocli"
	"github.com/percosis-labs/percosis/v16/x/valset-pref/client/queryproto"
	"github.com/percosis-labs/percosis/v16/x/valset-pref/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	cmd := percocli.QueryIndexCmd(types.ModuleName)
	cmd.AddCommand(GetCmdValSetPref())
	return cmd
}

// GetCmdValSetPref takes the  address and returns the existing validator set for that address.
func GetCmdValSetPref() *cobra.Command {
	return percocli.SimpleQueryCmd[*queryproto.UserValidatorPreferencesRequest](
		"val-set [address]",
		"Query the validator set for a specific user address", "",
		types.ModuleName, queryproto.NewQueryClient,
	)
}
