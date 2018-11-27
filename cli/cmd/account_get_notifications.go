package cmd

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	accountCmd.AddCommand(getAccountNotificationsCmd)
}

var getAccountNotificationsCmd = &cobra.Command{
	Use:   "notifications",
	Short: "Get Notifications",
	RunE: func(cmd *cobra.Command, args []string) error {
		if client.Token == "" {
			access, _, err := client.IdentityToken()
			if err != nil {
				fmt.Println(err)
				return err
			}
			client.Token = access.Token.ID
		}

		notifications, _, err := client.Notifications()
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(notifications)

		return nil
	},
}
