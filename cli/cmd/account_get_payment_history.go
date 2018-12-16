package cmd

import (
	"context"
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
		ctx := context.Background()
		if client.Token == "" {
			access, _, err := client.IdentityToken(ctx)
			if err != nil {
				fmt.Println(err)
				return err
			}
			client.Token = access.Token.ID
		}
		paymentHistory, _, err := client.PaymentHistory(ctx)
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(paymentHistory)

		return nil
	},
}
