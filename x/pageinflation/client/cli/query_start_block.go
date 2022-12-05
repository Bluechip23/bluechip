package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/smartdev0328/bluechip/x/pageinflation/types"
	"github.com/spf13/cobra"
)

func CmdShowStartBlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-start-block",
		Short: "shows startBlock",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetStartBlockRequest{}

			res, err := queryClient.StartBlock(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
