package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	flag "github.com/spf13/pflag"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/percosis-labs/percosis/osmoutils/percocli"
	"github.com/percosis-labs/percosis/v16/x/gamm/pool-models/balancer"
	"github.com/percosis-labs/percosis/v16/x/gamm/pool-models/stableswap"
	"github.com/percosis-labs/percosis/v16/x/poolmanager/types"
)

func NewTxCmd() *cobra.Command {
	txCmd := percocli.TxIndexCmd(types.ModuleName)

	percocli.AddTxCmd(txCmd, NewSwapExactAmountInCmd)
	percocli.AddTxCmd(txCmd, NewSwapExactAmountOutCmd)
	percocli.AddTxCmd(txCmd, NewSplitRouteSwapExactAmountIn)
	percocli.AddTxCmd(txCmd, NewSplitRouteSwapExactAmountOut)

	txCmd.AddCommand(
		NewCreatePoolCmd(),
	)

	return txCmd
}

func NewSwapExactAmountInCmd() (*percocli.TxCliDesc, *types.MsgSwapExactAmountIn) {
	return &percocli.TxCliDesc{
		Use:     "swap-exact-amount-in [token-in] [token-out-min-amount]",
		Short:   "swap exact amount in",
		Example: "percosisd tx poolmanager swap-exact-amount-in 2000000ufury 1 --swap-route-pool-ids 5 --swap-route-denoms uion --from val --keyring-backend test -b=block --chain-id=localpercosis --fees 10000ufury",
		CustomFieldParsers: map[string]percocli.CustomFieldParserFn{
			"Routes": percocli.FlagOnlyParser(swapAmountInRoutes),
		},
		Flags: percocli.FlagDesc{RequiredFlags: []*flag.FlagSet{FlagSetMultihopSwapRoutes()}},
	}, &types.MsgSwapExactAmountIn{}
}

func NewSwapExactAmountOutCmd() (*percocli.TxCliDesc, *types.MsgSwapExactAmountOut) {
	// Can't get rid of this parser without a break, because the args are out of order.
	return &percocli.TxCliDesc{
		Use:              "swap-exact-amount-out [token-out] [token-in-max-amount]",
		Short:            "swap exact amount out",
		Example:          "percosisd tx poolmanager swap-exact-amount-out 100uion 1000000 --swap-route-pool-ids 1 --swap-route-denoms ufury --from val --keyring-backend test -b=block --chain-id=localpercosis --fees 10000ufury",
		NumArgs:          2,
		ParseAndBuildMsg: NewBuildSwapExactAmountOutMsg,
		Flags:            percocli.FlagDesc{RequiredFlags: []*flag.FlagSet{FlagSetMultihopSwapRoutes()}},
	}, &types.MsgSwapExactAmountOut{}
}

func NewSplitRouteSwapExactAmountIn() (*percocli.TxCliDesc, *types.MsgSplitRouteSwapExactAmountIn) {
	return &percocli.TxCliDesc{
		Use:   "split-route-swap-exact-amount-in [token-in-denom] [token-out-min-amount] [flags]",
		Short: "split route swap exact amount in",
		Example: `percosisd tx poolmanager split-route-swap-exact-amount-in ufury 1 --routes-file="./routes.json" --from val --keyring-backend test -b=block --chain-id=localpercosis --fees 10000ufury
		- routes.json
		{
			"Route": [
			  {
			  "swap_amount_in_route": [
				{
				"pool_id": 1,
				"token_out_denom": "uion"
				},
				{
				"pool_id": 2,
				"token_out_denom": "ufury"
				}
			  ],
			  "token_in_amount": 1000
			  },
			  {
			  "swap_amount_in_route": [
				{
				"pool_id": 3,
				"token_out_denom": "bar"
				},
				{
				"pool_id": 4,
				"token_out_denom": "ufury"
				}
			  ],
			  "token_in_amount": 999
			  }
			]
		}
		`,
		CustomFieldParsers: map[string]percocli.CustomFieldParserFn{
			"Routes": percocli.FlagOnlyParser(NewMsgNewSplitRouteSwapExactAmountIn),
		},
		Flags: percocli.FlagDesc{
			RequiredFlags: []*flag.FlagSet{FlagSetCreateRoutes()},
		},
	}, &types.MsgSplitRouteSwapExactAmountIn{}
}

func NewSplitRouteSwapExactAmountOut() (*percocli.TxCliDesc, *types.MsgSplitRouteSwapExactAmountOut) {
	return &percocli.TxCliDesc{
		Use:   "split-route-swap-exact-amount-out [token-out-denom] [token-in-max-amount] [flags]",
		Short: "split route swap exact amount out",
		Example: `percosisd tx poolmanager split-route-swap-exact-amount-out ufury 1 --routes-file="./routes.json" --from val --keyring-backend test -b=block --chain-id=localpercosis --fees 10000ufury
		- routes.json
		{
			"route": [
				{
				"swap_amount_out_route": [
					{
					"pool_id": 1,
					"token_in_denom": "uion"
					},
					{
					"pool_id": 2,
					"token_in_denom": "ufury"
					}
				],
				"token_out_amount": 1000
				},
				{
				"swap_amount_out_route": [
					{
					"pool_id": 3,
					"token_in_denom": "uion"
					},
					{
					"pool_id": 4,
					"token_in_denom": "ufury"
					}
				],
				"token_out_amount": 999
				}
			]
			}
		`,
		CustomFieldParsers: map[string]percocli.CustomFieldParserFn{
			"Routes": percocli.FlagOnlyParser(NewMsgNewSplitRouteSwapExactAmountOut),
		},
		Flags: percocli.FlagDesc{
			RequiredFlags: []*flag.FlagSet{FlagSetCreateRoutes()},
		},
	}, &types.MsgSplitRouteSwapExactAmountOut{}
}

func NewMsgNewSplitRouteSwapExactAmountOut(fs *flag.FlagSet) ([]types.SwapAmountOutSplitRoute, error) {
	routesFile, _ := fs.GetString(FlagRoutesFile)
	if routesFile == "" {
		return nil, fmt.Errorf("must pass in a routes json using the --%s flag", FlagRoutesFile)
	}

	contents, err := os.ReadFile(routesFile)
	if err != nil {
		return nil, err
	}

	var splitRouteJSONdata RoutesOut
	err = json.Unmarshal(contents, &splitRouteJSONdata)
	if err != nil {
		return nil, err
	}

	var splitRouteProto []types.SwapAmountOutSplitRoute
	for _, route := range splitRouteJSONdata.Route {
		protoRoute := types.SwapAmountOutSplitRoute{
			TokenOutAmount: sdk.NewInt(route.TokenOutAmount),
		}
		protoRoute.Pools = append(protoRoute.Pools, route.Pools...)
		splitRouteProto = append(splitRouteProto, protoRoute)
	}

	return splitRouteProto, nil
}

func NewMsgNewSplitRouteSwapExactAmountIn(fs *flag.FlagSet) ([]types.SwapAmountInSplitRoute, error) {
	routesFile, _ := fs.GetString(FlagRoutesFile)
	if routesFile == "" {
		return nil, fmt.Errorf("must pass in a routes json using the --%s flag", FlagRoutesFile)
	}

	contents, err := os.ReadFile(routesFile)
	if err != nil {
		return nil, err
	}

	var splitRouteJSONdata RoutesIn
	err = json.Unmarshal(contents, &splitRouteJSONdata)
	if err != nil {
		return nil, err
	}

	var splitRouteProto []types.SwapAmountInSplitRoute
	for _, route := range splitRouteJSONdata.Route {
		protoRoute := types.SwapAmountInSplitRoute{
			TokenInAmount: sdk.NewInt(route.TokenInAmount),
		}
		protoRoute.Pools = append(protoRoute.Pools, route.Pools...)
		splitRouteProto = append(splitRouteProto, protoRoute)
	}

	return splitRouteProto, nil
}

func NewBuildSwapExactAmountOutMsg(clientCtx client.Context, args []string, fs *flag.FlagSet) (sdk.Msg, error) {
	tokenOutStr, tokenInMaxAmountStr := args[0], args[1]
	routes, err := swapAmountOutRoutes(fs)
	if err != nil {
		return nil, err
	}

	tokenOut, err := sdk.ParseCoinNormalized(tokenOutStr)
	if err != nil {
		return nil, err
	}

	tokenInMaxAmount, ok := sdk.NewIntFromString(tokenInMaxAmountStr)
	if !ok {
		return nil, errors.New("invalid token in max amount")
	}
	return &types.MsgSwapExactAmountOut{
		Sender:           clientCtx.GetFromAddress().String(),
		Routes:           routes,
		TokenInMaxAmount: tokenInMaxAmount,
		TokenOut:         tokenOut,
	}, nil
}

func NewCreatePoolCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-pool [flags]",
		Short: "create a new pool and provide the liquidity to it",
		Long:  `Must provide path to a pool JSON file (--pool-file) describing the pool to be created`,
		Example: `Sample pool JSON file contents:
{
	"weights": "4uatom,4perco,2uakt",
	"initial-deposit": "100uatom,5perco,20uakt",
	"swap-fee": "0.01",
	"exit-fee": "0.01",
	"future-governor": "168h"
}
`,
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			txf, msg, err := NewBuildCreateBalancerPoolMsg(clientCtx, txf, cmd.Flags())
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagSetCreatePool())
	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagPoolFile)

	return cmd
}

func NewBuildCreateBalancerPoolMsg(clientCtx client.Context, txf tx.Factory, fs *flag.FlagSet) (tx.Factory, sdk.Msg, error) {
	pool, err := parseCreateBalancerPoolFlags(fs)
	if err != nil {
		return txf, nil, fmt.Errorf("failed to parse pool: %w", err)
	}

	deposit, err := sdk.ParseCoinsNormalized(pool.InitialDeposit)
	if err != nil {
		return txf, nil, err
	}

	poolAssetCoins, err := sdk.ParseDecCoins(pool.Weights)
	if err != nil {
		return txf, nil, err
	}

	if len(deposit) != len(poolAssetCoins) {
		return txf, nil, errors.New("deposit tokens and token weights should have same length")
	}

	spreadFactor, err := sdk.NewDecFromStr(pool.SwapFee)
	if err != nil {
		return txf, nil, err
	}

	exitFee, err := sdk.NewDecFromStr(pool.ExitFee)
	if err != nil {
		return txf, nil, err
	}

	var poolAssets []balancer.PoolAsset
	for i := 0; i < len(poolAssetCoins); i++ {
		if poolAssetCoins[i].Denom != deposit[i].Denom {
			return txf, nil, errors.New("deposit tokens and token weights should have same denom order")
		}

		poolAssets = append(poolAssets, balancer.PoolAsset{
			Weight: poolAssetCoins[i].Amount.RoundInt(),
			Token:  deposit[i],
		})
	}

	poolParams := &balancer.PoolParams{
		SwapFee: spreadFactor,
		ExitFee: exitFee,
	}

	msg := &balancer.MsgCreateBalancerPool{
		Sender:             clientCtx.GetFromAddress().String(),
		PoolParams:         poolParams,
		PoolAssets:         poolAssets,
		FuturePoolGovernor: pool.FutureGovernor,
	}

	if (pool.SmoothWeightChangeParams != smoothWeightChangeParamsInputs{}) {
		duration, err := time.ParseDuration(pool.SmoothWeightChangeParams.Duration)
		if err != nil {
			return txf, nil, fmt.Errorf("could not parse duration: %w", err)
		}

		targetPoolAssetCoins, err := sdk.ParseDecCoins(pool.SmoothWeightChangeParams.TargetPoolWeights)
		if err != nil {
			return txf, nil, err
		}

		var targetPoolAssets []balancer.PoolAsset
		for i := 0; i < len(targetPoolAssetCoins); i++ {
			if targetPoolAssetCoins[i].Denom != poolAssetCoins[i].Denom {
				return txf, nil, errors.New("initial pool weights and target pool weights should have same denom order")
			}

			targetPoolAssets = append(targetPoolAssets, balancer.PoolAsset{
				Weight: targetPoolAssetCoins[i].Amount.RoundInt(),
				Token:  deposit[i],
				// TODO: This doesn't make sense. Should only use denom, not an sdk.Coin
			})
		}

		smoothWeightParams := balancer.SmoothWeightChangeParams{
			Duration:           duration,
			InitialPoolWeights: poolAssets,
			TargetPoolWeights:  targetPoolAssets,
		}

		if pool.SmoothWeightChangeParams.StartTime != "" {
			startTime, err := time.Parse(time.RFC3339, pool.SmoothWeightChangeParams.StartTime)
			if err != nil {
				return txf, nil, fmt.Errorf("could not parse time: %w", err)
			}

			smoothWeightParams.StartTime = startTime
		}

		msg.PoolParams.SmoothWeightChangeParams = &smoothWeightParams
	}

	return txf, msg, nil
}

// Apologies to whoever has to touch this next, this code is horrendous
func NewBuildCreateStableswapPoolMsg(clientCtx client.Context, txf tx.Factory, fs *flag.FlagSet) (tx.Factory, sdk.Msg, error) {
	flags, err := parseCreateStableswapPoolFlags(fs)
	if err != nil {
		return txf, nil, fmt.Errorf("failed to parse pool: %w", err)
	}

	deposit, err := ParseCoinsNoSort(flags.InitialDeposit)
	if err != nil {
		return txf, nil, err
	}

	spreadFactor, err := sdk.NewDecFromStr(flags.SwapFee)
	if err != nil {
		return txf, nil, err
	}

	exitFee, err := sdk.NewDecFromStr(flags.ExitFee)
	if err != nil {
		return txf, nil, err
	}

	poolParams := &stableswap.PoolParams{
		SwapFee: spreadFactor,
		ExitFee: exitFee,
	}

	scalingFactors := []uint64{}
	trimmedSfString := strings.Trim(flags.ScalingFactors, "[] {}")
	if len(trimmedSfString) > 0 {
		ints := strings.Split(trimmedSfString, ",")
		for _, i := range ints {
			u, err := strconv.ParseUint(i, 10, 64)
			if err != nil {
				return txf, nil, err
			}
			scalingFactors = append(scalingFactors, u)
		}
		if len(scalingFactors) != len(deposit) {
			return txf, nil, fmt.Errorf("number of scaling factors doesn't match number of assets")
		}
	}

	msg := &stableswap.MsgCreateStableswapPool{
		Sender:                  clientCtx.GetFromAddress().String(),
		PoolParams:              poolParams,
		InitialPoolLiquidity:    deposit,
		ScalingFactors:          scalingFactors,
		ScalingFactorController: flags.ScalingFactorController,
		FuturePoolGovernor:      flags.FutureGovernor,
	}

	return txf, msg, nil
}

// ParseCoinsNoSort parses coins from coinsStr but does not sort them.
// Returns error if parsing fails.
func ParseCoinsNoSort(coinsStr string) (sdk.Coins, error) {
	coinStrs := strings.Split(coinsStr, ",")
	decCoins := make(sdk.DecCoins, len(coinStrs))
	for i, coinStr := range coinStrs {
		coin, err := sdk.ParseDecCoin(coinStr)
		if err != nil {
			return sdk.Coins{}, err
		}

		decCoins[i] = coin
	}
	return sdk.NormalizeCoins(decCoins), nil
}
