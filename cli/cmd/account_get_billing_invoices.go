package cmd

import (
	"context"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	accountCmd.AddCommand(getAccountBillingInvoicesCmd)
}

var getAccountBillingInvoicesCmd = &cobra.Command{
	Use:   "invoices",
	Short: "Get Billing Invoices",
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
		invoices, _, err := client.BillingInvoices(ctx)
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(invoices)

		return nil
	},
}
