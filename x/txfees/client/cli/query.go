package cli

import (
	"github.com/spf13/cobra"

	"github.com/percosis-labs/percosis/osmoutils/percocli"
	"github.com/percosis-labs/percosis/v16/x/txfees/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	cmd := percocli.QueryIndexCmd(types.ModuleName)

	cmd.AddCommand(
		GetCmdFeeTokens(),
		GetCmdDenomPoolID(),
		GetCmdBaseDenom(),
	)

	return cmd
}

func GetCmdFeeTokens() *cobra.Command {
	return percocli.SimpleQueryCmd[*types.QueryFeeTokensRequest](
		"fee-tokens",
		"Query the list of non-basedenom fee tokens and their associated pool ids",
		`{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} fee-tokens
`,
		types.ModuleName, types.NewQueryClient,
	)
}

func GetCmdDenomPoolID() *cobra.Command {
	return percocli.SimpleQueryCmd[*types.QueryDenomPoolIdRequest](
		"denom-pool-id",
		"Query the pool id associated with a specific whitelisted fee token",
		`{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} denom-pool-id [denom]
`,
		types.ModuleName, types.NewQueryClient,
	)
}

func GetCmdBaseDenom() *cobra.Command {
	return percocli.SimpleQueryCmd[*types.QueryBaseDenomRequest](
		"base-denom",
		"Query the base fee denom",
		`{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} base-denom
`,
		types.ModuleName, types.NewQueryClient,
	)
}
