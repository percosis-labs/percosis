package cli

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/percosis-labs/percosis/osmoutils/percocli"

	"github.com/percosis-labs/percosis/v16/x/protorev/types"
)

// NewCmdQuery returns the cli query commands for this module
func NewCmdQuery() *cobra.Command {
	cmd := percocli.QueryIndexCmd(types.ModuleName)

	percocli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryParamsCmd)
	percocli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryNumberOfTradesCmd)
	percocli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryProfitsByDenomCmd)
	percocli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryAllProfitsCmd)
	percocli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryStatisticsByRouteCmd)
	percocli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryAllRouteStatisticsCmd)
	percocli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryTokenPairArbRoutesCmd)
	percocli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryAdminAccountCmd)
	percocli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryDeveloperAccountCmd)
	percocli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryMaxPoolPointsPerTxCmd)
	percocli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryMaxPoolPointsPerBlockCmd)
	percocli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryBaseDenomsCmd)
	percocli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryEnabledCmd)
	percocli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryPoolWeightsCmd)
	percocli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryPoolCmd)

	return cmd
}

// NewQueryParamsCmd returns the command to query the module params
func NewQueryParamsCmd() (*percocli.QueryDescriptor, *types.QueryParamsRequest) {
	return &percocli.QueryDescriptor{
		Use:   "params",
		Short: "Query the module params",
	}, &types.QueryParamsRequest{}
}

// NewQueryNumberOfTradesCmd returns the command to query the number of trades executed by protorev
func NewQueryNumberOfTradesCmd() (*percocli.QueryDescriptor, *types.QueryGetProtoRevNumberOfTradesRequest) {
	return &percocli.QueryDescriptor{
		Use:   "number-of-trades",
		Short: "Query the number of cyclic arbitrage trades protorev has executed",
	}, &types.QueryGetProtoRevNumberOfTradesRequest{}
}

// NewQueryProfitsByDenomCmd returns the command to query the profits of protorev by denom
func NewQueryProfitsByDenomCmd() (*percocli.QueryDescriptor, *types.QueryGetProtoRevProfitsByDenomRequest) {
	return &percocli.QueryDescriptor{
		Use:   "profits-by-denom [denom]",
		Short: "Query the profits of protorev by denom",
		Long:  `{{.Short}}{{.ExampleHeader}}{{.CommandPrefix}} profits-by-denom ufury`,
	}, &types.QueryGetProtoRevProfitsByDenomRequest{}
}

// NewQueryAllProfitsCmd returns the command to query all profits of protorev
func NewQueryAllProfitsCmd() (*percocli.QueryDescriptor, *types.QueryGetProtoRevAllProfitsRequest) {
	return &percocli.QueryDescriptor{
		Use:   "all-profits",
		Short: "Query all ProtoRev profits",
	}, &types.QueryGetProtoRevAllProfitsRequest{}
}

// NewQueryStatisticsByRoute returns the command to query the statistics of protorev by route
func NewQueryStatisticsByRouteCmd() (*percocli.QueryDescriptor, *types.QueryGetProtoRevStatisticsByRouteRequest) {
	return &percocli.QueryDescriptor{
		Use:                "statistics-by-route [route]",
		Short:              "Query statistics about a specific arbitrage route",
		Long:               `{{.Short}}{{.ExampleHeader}}{{.CommandPrefix}} statistics-by-route [1,2,3]`,
		CustomFieldParsers: map[string]percocli.CustomFieldParserFn{"Route": parseRoute},
	}, &types.QueryGetProtoRevStatisticsByRouteRequest{}
}

// NewQueryAllRouteStatisticsCmd returns the command to query all route statistics of protorev
func NewQueryAllRouteStatisticsCmd() (*percocli.QueryDescriptor, *types.QueryGetProtoRevAllRouteStatisticsRequest) {
	return &percocli.QueryDescriptor{
		Use:   "all-statistics",
		Short: "Query all ProtoRev statistics",
	}, &types.QueryGetProtoRevAllRouteStatisticsRequest{}
}

// NewQueryTokenPairArbRoutesCmd returns the command to query the token pair arb routes
func NewQueryTokenPairArbRoutesCmd() (*percocli.QueryDescriptor, *types.QueryGetProtoRevTokenPairArbRoutesRequest) {
	return &percocli.QueryDescriptor{
		Use:   "hot-routes",
		Short: "Query the ProtoRev hot routes currently being used",
	}, &types.QueryGetProtoRevTokenPairArbRoutesRequest{}
}

// NewQueryAdminAccountCmd returns the command to query the admin account
func NewQueryAdminAccountCmd() (*percocli.QueryDescriptor, *types.QueryGetProtoRevAdminAccountRequest) {
	return &percocli.QueryDescriptor{
		Use:   "admin-account",
		Short: "Query the admin account",
	}, &types.QueryGetProtoRevAdminAccountRequest{}
}

// NewQueryDeveloperAccountCmd returns the command to query the developer account
func NewQueryDeveloperAccountCmd() (*percocli.QueryDescriptor, *types.QueryGetProtoRevDeveloperAccountRequest) {
	return &percocli.QueryDescriptor{
		Use:   "developer-account",
		Short: "Query the developer account",
	}, &types.QueryGetProtoRevDeveloperAccountRequest{}
}

// NewQueryMaxPoolPointsPerTxCmd returns the command to query the max pool points per tx
func NewQueryMaxPoolPointsPerTxCmd() (*percocli.QueryDescriptor, *types.QueryGetProtoRevMaxPoolPointsPerTxRequest) {
	return &percocli.QueryDescriptor{
		Use:   "max-pool-points-per-tx",
		Short: "Query the max pool points per tx",
	}, &types.QueryGetProtoRevMaxPoolPointsPerTxRequest{}
}

// NewQueryMaxPoolPointsPerBlockCmd returns the command to query the max pool points per block
func NewQueryMaxPoolPointsPerBlockCmd() (*percocli.QueryDescriptor, *types.QueryGetProtoRevMaxPoolPointsPerBlockRequest) {
	return &percocli.QueryDescriptor{
		Use:   "max-pool-points-per-block",
		Short: "Query the max pool points per block",
	}, &types.QueryGetProtoRevMaxPoolPointsPerBlockRequest{}
}

// NewQueryBaseDenomsCmd returns the command to query the base denoms
func NewQueryBaseDenomsCmd() (*percocli.QueryDescriptor, *types.QueryGetProtoRevBaseDenomsRequest) {
	return &percocli.QueryDescriptor{
		Use:   "base-denoms",
		Short: "Query the base denoms used to construct arbitrage routes",
	}, &types.QueryGetProtoRevBaseDenomsRequest{}
}

// NewQueryEnabled returns the command to query the enabled status of protorev
func NewQueryEnabledCmd() (*percocli.QueryDescriptor, *types.QueryGetProtoRevEnabledRequest) {
	return &percocli.QueryDescriptor{
		Use:   "enabled",
		Short: "Query whether protorev is currently enabled",
	}, &types.QueryGetProtoRevEnabledRequest{}
}

// NewQueryPoolWeightsCmd returns the command to query the pool weights of protorev
func NewQueryPoolWeightsCmd() (*percocli.QueryDescriptor, *types.QueryGetProtoRevPoolWeightsRequest) {
	return &percocli.QueryDescriptor{
		Use:   "pool-weights",
		Short: "Query the pool weights used to determine how computationally expensive a route is",
	}, &types.QueryGetProtoRevPoolWeightsRequest{}
}

// NewQueryPoolCmd returns the command to query the pool id for a given denom pair stored via the highest liquidity method in ProtoRev
func NewQueryPoolCmd() (*percocli.QueryDescriptor, *types.QueryGetProtoRevPoolRequest) {
	return &percocli.QueryDescriptor{
		Use:   "pool [base_denom] [other_denom]",
		Short: "Query the pool id for a given denom pair stored via the highest liquidity method in ProtoRev",
	}, &types.QueryGetProtoRevPoolRequest{}
}

// convert a string array "[1,2,3]" to []uint64
func parseRoute(arg string, _ *pflag.FlagSet) (any, percocli.FieldReadLocation, error) {
	var route []uint64
	err := json.Unmarshal([]byte(arg), &route)
	if err != nil {
		return nil, percocli.UsedArg, err
	}
	return route, percocli.UsedArg, err
}
