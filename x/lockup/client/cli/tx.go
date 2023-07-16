package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/percosis-labs/percosis/osmoutils/percocli"
	"github.com/percosis-labs/percosis/v16/x/lockup/types"
)

// GetTxCmd returns the transaction commands for this module.
func GetTxCmd() *cobra.Command {
	cmd := percocli.TxIndexCmd(types.ModuleName)
	percocli.AddTxCmd(cmd, NewLockTokensCmd)
	percocli.AddTxCmd(cmd, NewBeginUnlockingAllCmd)
	percocli.AddTxCmd(cmd, NewBeginUnlockByIDCmd)
	percocli.AddTxCmd(cmd, NewForceUnlockByIdCmd)
	percocli.AddTxCmd(cmd, NewSetRewardReceiverAddress)

	return cmd
}

func NewLockTokensCmd() (*percocli.TxCliDesc, *types.MsgLockTokens) {
	return &percocli.TxCliDesc{
		Use:   "lock-tokens [tokens]",
		Short: "lock tokens into lockup pool from user account",
		CustomFlagOverrides: map[string]string{
			"duration": FlagDuration,
		},
		Flags: percocli.FlagDesc{RequiredFlags: []*pflag.FlagSet{FlagSetLockTokens()}},
	}, &types.MsgLockTokens{}
}

// TODO: We should change the Use string to be unlock-all
func NewBeginUnlockingAllCmd() (*percocli.TxCliDesc, *types.MsgBeginUnlockingAll) {
	return &percocli.TxCliDesc{
		Use:   "begin-unlock-tokens",
		Short: "begin unlock not unlocking tokens from lockup pool for sender",
	}, &types.MsgBeginUnlockingAll{}
}

// NewBeginUnlockByIDCmd unlocks individual period lock by ID.
func NewBeginUnlockByIDCmd() (*percocli.TxCliDesc, *types.MsgBeginUnlocking) {
	return &percocli.TxCliDesc{
		Use:   "begin-unlock-by-id [id]",
		Short: "begin unlock individual period lock by ID",
		CustomFlagOverrides: map[string]string{
			"coins": FlagAmount,
		},
		Flags: percocli.FlagDesc{OptionalFlags: []*pflag.FlagSet{FlagSetUnlockTokens()}},
	}, &types.MsgBeginUnlocking{}
}

// NewForceUnlockByIdCmd force unlocks individual period lock by ID if proper permissions exist.
func NewForceUnlockByIdCmd() (*percocli.TxCliDesc, *types.MsgForceUnlock) {
	return &percocli.TxCliDesc{
		Use:   "force-unlock-by-id [id]",
		Short: "force unlocks individual period lock by ID",
		Long:  "force unlocks individual period lock by ID. if no amount provided, entire lock is unlocked",
		CustomFlagOverrides: map[string]string{
			"coins": FlagAmount,
		},
		Flags: percocli.FlagDesc{OptionalFlags: []*pflag.FlagSet{FlagSetUnlockTokens()}},
	}, &types.MsgForceUnlock{}
}

// NewSetRewardReceiverAddress sets the reward receiver address.
func NewSetRewardReceiverAddress() (*percocli.TxCliDesc, *types.MsgSetRewardReceiverAddress) {
	return &percocli.TxCliDesc{
		Use:   "set-reward-receiver-address [lock-id] [reward-receiver]",
		Short: "sets reward receiver address for the designated lock id",
		Long:  "sets reward receiver address for the designated lock id",
	}, &types.MsgSetRewardReceiverAddress{}
}
