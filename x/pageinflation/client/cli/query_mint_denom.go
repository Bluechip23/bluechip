package cli

import (
	"context"

	"github.com/Smartdev0328/bluechip/x/pageinflation/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdShowMintDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-mint-denom",
		Short: "shows mintDenom",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetMintDenomRequest{}

			res, err := queryClient.MintDenom(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
