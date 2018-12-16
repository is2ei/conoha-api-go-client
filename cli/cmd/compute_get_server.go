package cmd

import (
	"context"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	getComputeServerCmd.Flags().StringVarP(&serverID, "serverId", "i", "", "Server ID")
	serverCmd.AddCommand(getComputeServerCmd)
}

var getComputeServerCmd = &cobra.Command{
	Use:   "detail",
	Short: "Get server detail",
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

		server, _, err := client.ComputeServer(ctx, serverID)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if verbose {
			pp.Println(server)
		} else {
			fmt.Printf("name[%s%s%s], id[%s%s%s], status[%s%s%s], ",
				green, server.Name, normal, yellow, server.ID, normal, green, server.Status, normal)
			for _, address := range server.Addresses {
				for _, a := range address {
					fmt.Printf("address[%s%s%s], ", yellow, a.Addr, normal)
				}
				fmt.Printf("\n")
			}
		}

		return nil
	},
}
