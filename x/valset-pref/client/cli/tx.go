package valsetprefcli

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/percosis-labs/percosis/osmoutils"
	"github.com/percosis-labs/percosis/osmoutils/percocli"
	"github.com/percosis-labs/percosis/v16/x/valset-pref/types"
)

func GetTxCmd() *cobra.Command {
	txCmd := percocli.TxIndexCmd(types.ModuleName)
	percocli.AddTxCmd(txCmd, NewSetValSetCmd)
	percocli.AddTxCmd(txCmd, NewDelValSetCmd)
	percocli.AddTxCmd(txCmd, NewUnDelValSetCmd)
	percocli.AddTxCmd(txCmd, NewReDelValSetCmd)
	percocli.AddTxCmd(txCmd, NewWithRewValSetCmd)
	return txCmd
}

func NewSetValSetCmd() (*percocli.TxCliDesc, *types.MsgSetValidatorSetPreference) {
	return &percocli.TxCliDesc{
		Use:              "set-valset [delegator_addr] [validators] [weights]",
		Short:            "Creates a new validator set for the delegator with valOperAddress and weight",
		Example:          "percosisd tx valset-pref set-valset perco1... percovaloper1abc...,percovaloper1def...  0.56,0.44",
		NumArgs:          3,
		ParseAndBuildMsg: NewMsgSetValidatorSetPreference,
	}, &types.MsgSetValidatorSetPreference{}
}

func NewDelValSetCmd() (*percocli.TxCliDesc, *types.MsgDelegateToValidatorSet) {
	return &percocli.TxCliDesc{
		Use:     "delegate-valset [delegator_addr] [amount]",
		Short:   "Delegate tokens to existing valset using delegatorAddress and tokenAmount.",
		Example: "percosisd tx valset-pref delegate-valset perco1... 100stake",
		NumArgs: 2,
	}, &types.MsgDelegateToValidatorSet{}
}

func NewUnDelValSetCmd() (*percocli.TxCliDesc, *types.MsgUndelegateFromValidatorSet) {
	return &percocli.TxCliDesc{
		Use:     "undelegate-valset [delegator_addr] [amount]",
		Short:   "UnDelegate tokens from existing valset using delegatorAddress and tokenAmount.",
		Example: "percosisd tx valset-pref undelegate-valset perco1... 100stake",
		NumArgs: 2,
	}, &types.MsgUndelegateFromValidatorSet{}
}

func NewReDelValSetCmd() (*percocli.TxCliDesc, *types.MsgRedelegateValidatorSet) {
	return &percocli.TxCliDesc{
		Use:              "redelegate-valset [delegator_addr] [validators] [weights]",
		Short:            "Redelegate tokens from existing validators to new sets of validators",
		Example:          "percosisd tx valset-pref redelegate-valset  perco1... percovaloper1efg...,percovaloper1hij...  0.56,0.44",
		NumArgs:          3,
		ParseAndBuildMsg: NewMsgReDelValidatorSetPreference,
	}, &types.MsgRedelegateValidatorSet{}
}

func NewWithRewValSetCmd() (*percocli.TxCliDesc, *types.MsgWithdrawDelegationRewards) {
	return &percocli.TxCliDesc{
		Use:     "withdraw-reward-valset [delegator_addr]",
		Short:   "Withdraw delegation reward form the new validator set.",
		Example: "percosisd tx valset-pref withdraw-valset perco1...",
		NumArgs: 1,
	}, &types.MsgWithdrawDelegationRewards{}
}

func NewMsgSetValidatorSetPreference(clientCtx client.Context, args []string, fs *pflag.FlagSet) (sdk.Msg, error) {
	delAddr, err := sdk.AccAddressFromBech32(args[0])
	if err != nil {
		return nil, err
	}

	valset, err := ValidateValAddrAndWeight(args)
	if err != nil {
		return nil, err
	}

	return types.NewMsgSetValidatorSetPreference(
		delAddr,
		valset,
	), nil
}

func NewMsgReDelValidatorSetPreference(clientCtx client.Context, args []string, fs *pflag.FlagSet) (sdk.Msg, error) {
	delAddr, err := sdk.AccAddressFromBech32(args[0])
	if err != nil {
		return nil, err
	}

	valset, err := ValidateValAddrAndWeight(args)
	if err != nil {
		return nil, err
	}

	return types.NewMsgRedelegateValidatorSet(
		delAddr,
		valset,
	), nil
}

func ValidateValAddrAndWeight(args []string) ([]types.ValidatorPreference, error) {
	var valAddrs []string
	valAddrs = append(valAddrs, strings.Split(args[1], ",")...)

	weights, err := osmoutils.ParseSdkDecFromString(args[2], ",")
	if err != nil {
		return nil, err
	}

	if len(valAddrs) != len(weights) {
		return nil, fmt.Errorf("the length of validator addresses and weights not matched")
	}

	if len(valAddrs) == 0 {
		return nil, fmt.Errorf("records is empty")
	}

	var valset []types.ValidatorPreference
	for i, val := range valAddrs {
		valset = append(valset, types.ValidatorPreference{
			ValOperAddress: val,
			Weight:         weights[i],
		})
	}

	return valset, nil
}
