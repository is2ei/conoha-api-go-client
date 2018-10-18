package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serverId string

func init() {
	deleteComputeServerCmd.Flags().StringVarP(&serverId, "serverId", "i", "", "Serer ID")
	serverCmd.AddCommand(deleteComputeServerCmd)
}

var deleteComputeServerCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete server",
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
		_, err := client.DeleteComputeServer(serverId)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	},
}
