package cmd

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

var invoiceID string

func init() {
	getAccountBillingInvoiceCmd.Flags().StringVarP(&invoiceID, "id", "i", "", "Invoice ID")
	accountCmd.AddCommand(getAccountBillingInvoiceCmd)
}

var getAccountBillingInvoiceCmd = &cobra.Command{
	Use:   "invoice",
	Short: "Get Billing invoice",
	RunE: func(cmd *cobra.Command, args []string) error {
		if client.Token == "" {
			access, _, err := client.IdentityToken()
			if err != nil {
				fmt.Println(err)
				return err
			}
			client.Token = access.Token.ID
		}
		invoice, _, err := client.BillingInvoice(invoiceID)
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(invoice)

		return nil
	},
}
