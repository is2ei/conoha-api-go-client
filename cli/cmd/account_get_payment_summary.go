package cmd

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	accountCmd.AddCommand(getAccountPaymentSummaryCmd)
}

var getAccountPaymentSummaryCmd = &cobra.Command{
	Use:   "payment-summary",
	Short: "Get Payment Summary",
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
		paymentSummary, _, err := client.PaymentSummary()
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(paymentSummary)

		return nil
	},
}
