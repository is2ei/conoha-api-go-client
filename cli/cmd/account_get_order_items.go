package cmd

import (
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
		if client.Token == "" {
			access, _, err := client.IdentityToken()
			if err != nil {
				fmt.Println(err)
				return err
			}
			client.Token = access.Token.Id
		}
		items, _, err := client.OrderItems()
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(items)

		return nil
	},
}
