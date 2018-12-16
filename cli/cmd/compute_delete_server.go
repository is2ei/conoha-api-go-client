package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	deleteComputeServerCmd.Flags().StringVarP(&serverID, "serverId", "i", "", "Server ID")
	serverCmd.AddCommand(deleteComputeServerCmd)
}

var deleteComputeServerCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete server",
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
		_, err := client.DeleteComputeServer(ctx, serverID)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	},
}
