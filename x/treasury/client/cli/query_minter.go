package cli

import (
    "context"
    "strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
    "github.com/platine-network/platine/x/treasury/types"
)

func CmdListMinter() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-minter",
		Short: "list all minter",
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            pageReq, err := client.ReadPageRequest(cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryAllMinterRequest{
                Pagination: pageReq,
            }

            res, err := queryClient.MinterAll(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}

func CmdShowMinter() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-minter [id]",
		Short: "shows a minter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            queryClient := types.NewQueryClient(clientCtx)

            id, err := strconv.ParseUint(args[0], 10, 64)
            if err != nil {
                return err
            }

            params := &types.QueryGetMinterRequest{
                Id: id,
            }

            res, err := queryClient.Minter(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}
