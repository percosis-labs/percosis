package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/percosis-labs/percosis/osmoutils/percocli"
	"github.com/percosis-labs/percosis/v16/x/tokenfactory/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := percocli.TxIndexCmd(types.ModuleName)
	cmd.AddCommand(
		NewCreateDenomCmd(),
		NewMintCmd(),
		NewBurnCmd(),
		// NewForceTransferCmd(),
		NewChangeAdminCmd(),
		NewSetBeforeSendHookCmd(),
	)

	return cmd
}

func NewCreateDenomCmd() *cobra.Command {
	return percocli.BuildTxCli[*types.MsgCreateDenom](&percocli.TxCliDesc{
		Use:   "create-denom [subdenom] [flags]",
		Short: "create a new denom from an account. (Costs perco though!)",
	})
}

func NewMintCmd() *cobra.Command {
	return percocli.BuildTxCli[*types.MsgMint](&percocli.TxCliDesc{
		Use:   "mint [amount] [flags]",
		Short: "Mint a denom to an address. Must have admin authority to do so.",
	})
}

func NewBurnCmd() *cobra.Command {
	return percocli.BuildTxCli[*types.MsgBurn](&percocli.TxCliDesc{
		Use:   "burn [amount] [flags]",
		Short: "Burn tokens from an address. Must have admin authority to do so.",
	})
}

func NewChangeAdminCmd() *cobra.Command {
	return percocli.BuildTxCli[*types.MsgChangeAdmin](&percocli.TxCliDesc{
		Use:   "change-admin [denom] [new-admin-address] [flags]",
		Short: "Changes the admin address for a factory-created denom. Must have admin authority to do so.",
	})
}

// NewChangeAdminCmd broadcast MsgChangeAdmin
func NewSetBeforeSendHookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-beforesend-hook [denom] [cosmwasm-address] [flags]",
		Short: "Set a cosmwasm contract to be the beforesend hook for a factory-created denom. Must have admin authority to do so.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			msg := types.NewMsgSetBeforeSendHook(
				clientCtx.GetFromAddress().String(),
				args[0],
				args[1],
			)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
