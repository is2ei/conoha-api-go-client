package cmd

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

var itemID string

func init() {
	getAccountOrderItemCmd.Flags().StringVarP(&itemID, "item-id", "i", "", "Item ID")
	accountCmd.AddCommand(getAccountOrderItemCmd)
}

var getAccountOrderItemCmd = &cobra.Command{
	Use:   "item",
	Short: "Get order item",
	RunE: func(cmd *cobra.Command, args []string) error {
		if client.Token == "" {
			access, _, err := client.IdentityToken()
			if err != nil {
				fmt.Println(err)
				return err
			}
			client.Token = access.Token.ID
		}
		item, _, err := client.OrderItem(itemID)
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(item)

		return nil
	},
}
