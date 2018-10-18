package cmd

import (
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
		if client.Token == "" {
			access, _, err := client.IdentityToken()
			if err != nil {
				fmt.Println(err)
				return err
			} else {
				client.Token = access.Token.Id
			}
		}
		invoices, _, err := client.BillingInvoices()
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(invoices)

		return nil
	},
}
