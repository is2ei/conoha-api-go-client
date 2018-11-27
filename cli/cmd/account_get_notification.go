package cmd

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

var notificationCode string

func init() {
	getAccountNotificationCmd.Flags().StringVarP(&notificationCode, "code", "c", "", "Notification Code")
	accountCmd.AddCommand(getAccountNotificationCmd)
}

var getAccountNotificationCmd = &cobra.Command{
	Use:   "notification",
	Short: "Get Notification Detail",
	RunE: func(cmd *cobra.Command, args []string) error {
		if client.Token == "" {
			access, _, err := client.IdentityToken()
			if err != nil {
				fmt.Println(err)
				return err
			}
			client.Token = access.Token.ID
		}

		notification, _, err := client.Notification(notificationCode)
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(notification)

		return nil
	},
}
