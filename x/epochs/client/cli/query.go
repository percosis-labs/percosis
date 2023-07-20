package cli

import (
	"github.com/spf13/cobra"

	"github.com/osmosis-labs/osmosis/osmoutils/osmocli"
	"github.com/osmosis-labs/osmosis/x/epochs/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	cmd := percocli.QueryIndexCmd(types.ModuleName)
	percocli.AddQueryCmd(cmd, types.NewQueryClient, GetCmdEpochInfos)
	percocli.AddQueryCmd(cmd, types.NewQueryClient, GetCmdCurrentEpoch)

	return cmd
}

func GetCmdEpochInfos() (*percocli.QueryDescriptor, *types.QueryEpochsInfoRequest) {
	return &percocli.QueryDescriptor{
		Use:   "epoch-infos",
		Short: "Query running epoch infos.",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}}`,
		QueryFnName: "EpochInfos",
	}, &types.QueryEpochsInfoRequest{}
}

func GetCmdCurrentEpoch() (*percocli.QueryDescriptor, *types.QueryCurrentEpochRequest) {
	return &percocli.QueryDescriptor{
		Use:   "current-epoch",
		Short: "Query current epoch by specified identifier.",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} day`,
	}, &types.QueryCurrentEpochRequest{}
}
