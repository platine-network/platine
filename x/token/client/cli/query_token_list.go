package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/platine-network/platine/x/token/types"
	"github.com/spf13/cobra"
)

func CmdGetTokenList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token-list [owner]",
		Short: "list all token for the given owner address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			owner := args[0]
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryTokenListRequest{
				Owner: owner,
			}

			res, err := queryClient.TokenList(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

