package cli

import (
	"github.com/spf13/cobra"

	"github.com/percosis-labs/percosis/osmoutils/percocli"
	"github.com/percosis-labs/percosis/v16/x/cosmwasmpool/client/queryproto"
	"github.com/percosis-labs/percosis/v16/x/cosmwasmpool/types"
)

// NewQueryCmd returns the cli query commands for this module.
func NewQueryCmd() *cobra.Command {
	cmd := percocli.QueryIndexCmd(types.ModuleName)

	qcGetter := queryproto.NewQueryClient
	percocli.AddQueryCmd(cmd, qcGetter, GetCmdPools)
	percocli.AddQueryCmd(cmd, qcGetter, GetCmdContractInfoByPoolId)
	cmd.AddCommand(
		percocli.GetParams[*queryproto.ParamsRequest](
			types.ModuleName, queryproto.NewQueryClient),
	)

	return cmd
}

func GetCmdPools() (*percocli.QueryDescriptor, *queryproto.PoolsRequest) {
	return &percocli.QueryDescriptor{
		Use:   "pools",
		Short: "Query pools",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} pools`,
	}, &queryproto.PoolsRequest{}
}

func GetCmdContractInfoByPoolId() (*percocli.QueryDescriptor, *queryproto.ContractInfoByPoolIdRequest) {
	return &percocli.QueryDescriptor{
		Use:   "contract-info [pool-id]",
		Short: "Query contract info by pool id",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} pools`,
	}, &queryproto.ContractInfoByPoolIdRequest{}
}
