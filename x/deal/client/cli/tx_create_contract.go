package cli

import (
	"strconv"

	"github.com/Harry-027/deal/x/deal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-contract [deal-id] [consumer] [desc] [owner-eta] [expiry]",
		Short: "Broadcast message createContract",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDealId := args[0]
			argConsumer := args[1]
			argDesc := args[2]
			argOwnerETA := args[3]
			argExpiry := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateContract(
				clientCtx.GetFromAddress().String(),
				argDealId,
				argConsumer,
				argDesc,
				argOwnerETA,
				argExpiry,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
