package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/platine-network/platine/x/token/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-token [owner] [name] [symbol] [supply] [decimal] [mintable] [burnable]",
		Short: "Broadcast message create-token",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOwner := args[0]
			argName := args[1]
			argSymbol := args[2]
			argSupply, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			argDecimal, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}
			argMintable, err := cast.ToBoolE(args[5])
			if err != nil {
				return err
			}
			argBurnable, err := cast.ToBoolE(args[6])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateToken(
				clientCtx.GetFromAddress().String(),
				argOwner,
				argName,
				argSymbol,
				argSupply,
				argDecimal,
				argMintable,
				argBurnable,
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
