package cmd

import (
	"context"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	accountCmd.AddCommand(getAccountOrderItemsCmd)
}

var getAccountOrderItemsCmd = &cobra.Command{
	Use:   "items",
	Short: "Get order items",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		if client.Token == "" {
			access, _, err := client.IdentityToken(ctx)
			if err != nil {
				fmt.Println(err)
				return err
			}
			client.Token = access.Token.ID
		}
		items, _, err := client.OrderItems(ctx)
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(items)

		return nil
	},
}
