package cli

import (
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	"github.com/percosis-labs/percosis/osmoutils/percocli"
	"github.com/percosis-labs/percosis/v16/x/concentrated-liquidity/client/queryproto"
	"github.com/percosis-labs/percosis/v16/x/concentrated-liquidity/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	cmd := percocli.QueryIndexCmd(types.ModuleName)
	percocli.AddQueryCmd(cmd, queryproto.NewQueryClient, GetCmdPools)
	percocli.AddQueryCmd(cmd, queryproto.NewQueryClient, GetUserPositions)
	percocli.AddQueryCmd(cmd, queryproto.NewQueryClient, GetPositionById)
	percocli.AddQueryCmd(cmd, queryproto.NewQueryClient, GetClaimableSpreadRewards)
	percocli.AddQueryCmd(cmd, queryproto.NewQueryClient, GetClaimableIncentives)
	percocli.AddQueryCmd(cmd, queryproto.NewQueryClient, GetIncentiveRecords)
	percocli.AddQueryCmd(cmd, queryproto.NewQueryClient, GetCFMMPoolIdLinkFromConcentratedPoolId)
	percocli.AddQueryCmd(cmd, queryproto.NewQueryClient, GetTickLiquidityNetInDirection)
	percocli.AddQueryCmd(cmd, queryproto.NewQueryClient, GetPoolAccumulatorRewards)
	percocli.AddQueryCmd(cmd, queryproto.NewQueryClient, GetTickAccumulatorTrackers)
	cmd.AddCommand(
		percocli.GetParams[*queryproto.ParamsRequest](
			types.ModuleName, queryproto.NewQueryClient),
	)
	return cmd
}

func GetUserPositions() (*percocli.QueryDescriptor, *queryproto.UserPositionsRequest) {
	return &percocli.QueryDescriptor{
			Use:   "user-positions [address]",
			Short: "Query user's positions",
			Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} user-positions perco12smx2wdlyttvyzvzg54y2vnqwq2qjateuf7thj`,
			Flags:               percocli.FlagDesc{OptionalFlags: []*flag.FlagSet{FlagSetJustPoolId()}},
			CustomFlagOverrides: poolIdFlagOverride,
		},
		&queryproto.UserPositionsRequest{}
}

func GetPositionById() (*percocli.QueryDescriptor, *queryproto.PositionByIdRequest) {
	return &percocli.QueryDescriptor{
			Use:   "position-by-id [positionID]",
			Short: "Query position by ID",
			Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} position-by-id 53`,
		},
		&queryproto.PositionByIdRequest{}
}

func GetCmdPools() (*percocli.QueryDescriptor, *queryproto.PoolsRequest) {
	return &percocli.QueryDescriptor{
		Use:   "pools",
		Short: "Query pools",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} pools`,
	}, &queryproto.PoolsRequest{}
}

func GetClaimableSpreadRewards() (*percocli.QueryDescriptor, *queryproto.ClaimableSpreadRewardsRequest) {
	return &percocli.QueryDescriptor{
		Use:   "claimable-spread-rewards [positionID]",
		Short: "Query claimable spread rewards",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} claimable-spread-rewards 53`,
	}, &queryproto.ClaimableSpreadRewardsRequest{}
}

func GetClaimableIncentives() (*percocli.QueryDescriptor, *queryproto.ClaimableIncentivesRequest) {
	return &percocli.QueryDescriptor{
		Use:   "claimable-incentives [positionID]",
		Short: "Query claimable incentives",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} claimable-incentives 53`,
	}, &queryproto.ClaimableIncentivesRequest{}
}

func GetIncentiveRecords() (*percocli.QueryDescriptor, *queryproto.IncentiveRecordsRequest) {
	return &percocli.QueryDescriptor{
		Use:   "incentive-records [poolId]",
		Short: "Query incentive records for a given pool",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} incentive-records 1`,
	}, &queryproto.IncentiveRecordsRequest{}
}

func GetCFMMPoolIdLinkFromConcentratedPoolId() (*percocli.QueryDescriptor, *queryproto.CFMMPoolIdLinkFromConcentratedPoolIdRequest) {
	return &percocli.QueryDescriptor{
		Use:   "cfmm-pool-link-from-cl [poolId]",
		Short: "Query cfmm pool id link from concentrated pool id",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} cfmm-pool-link-from-cl 1`,
	}, &queryproto.CFMMPoolIdLinkFromConcentratedPoolIdRequest{}
}

func GetTotalLiquidity() (*percocli.QueryDescriptor, *queryproto.GetTotalLiquidityRequest) {
	return &percocli.QueryDescriptor{
		Use:   "total-liquidity",
		Short: "Query total liquidity across all concentrated pool",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} total-liquidity 1`,
	}, &queryproto.GetTotalLiquidityRequest{}
}

func GetTickLiquidityNetInDirection() (*percocli.QueryDescriptor, *queryproto.LiquidityNetInDirectionRequest) {
	return &percocli.QueryDescriptor{
		Use:   "liquidity-net-in-direction [pool-id] [token-in-denom] [start-tick] [use-current-tick] [bound-tick] [use-no-bound]",
		Short: "Query liquidity net in direction",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} 4 ufury "[-18000000]" true "[-9000000]" true`,
	}, &queryproto.LiquidityNetInDirectionRequest{}
}

func GetPoolAccumulatorRewards() (*percocli.QueryDescriptor, *queryproto.PoolAccumulatorRewardsRequest) {
	return &percocli.QueryDescriptor{
		Use:   "pool-accumulator-rewards [pool-id]",
		Short: "Query pool accumulator rewards",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} pool-accumulator-rewards 1`,
	}, &queryproto.PoolAccumulatorRewardsRequest{}
}

func GetTickAccumulatorTrackers() (*percocli.QueryDescriptor, *queryproto.TickAccumulatorTrackersRequest) {
	return &percocli.QueryDescriptor{
		Use:   "tick-accumulator-trackers [pool-id] [tick-index]",
		Short: "Query tick accumulator trackers",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} tick-accumulator-trackers 1 "[-18000000]"`,
	}, &queryproto.TickAccumulatorTrackersRequest{}
}
