package cmd

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	accountCmd.AddCommand(getAccountPaymentHistoryCmd)
}

var getAccountPaymentHistoryCmd = &cobra.Command{
	Use:   "payment-history",
	Short: "Get Payment History",
	RunE: func(cmd *cobra.Command, args []string) error {
		if client.Token == "" {
			access, _, err := client.IdentityToken()
			if err != nil {
				fmt.Println(err)
				return err
			}
			client.Token = access.Token.ID
		}
		paymentHistory, _, err := client.PaymentHistory()
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(paymentHistory)

		return nil
	},
}
