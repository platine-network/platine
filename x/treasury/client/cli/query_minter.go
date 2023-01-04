package cli

import (
	"context"
	"fmt"

	"github.com/platine-network/platine/x/treasury/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

func CmdQueryEpochProvision() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "epoch-provision",
		Short: "Query the current treasury epoch provision",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.EpochProvision(context.Background(), &types.QueryEpochProvisionRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintString(fmt.Sprintf("%s\n", &res.EpochProvision))
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
