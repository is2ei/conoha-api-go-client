package cmd

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	getComputeServerCmd.Flags().StringVarP(&serverId, "serverId", "i", "", "Server ID")
	serverCmd.AddCommand(getComputeServerCmd)
}

var getComputeServerCmd = &cobra.Command{
	Use:   "detail",
	Short: "Get server detail",
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

		server, _, err := client.ComputeServer(serverId)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if verbose {
			pp.Println(server)
		} else {
			fmt.Printf("name[%s%s%s], id[%s%s%s], status[%s%s%s], ",
				green, server.Name, normal, yellow, server.Id, normal, green, server.Status, normal)
			for _, address := range server.Addresses {
				for _, a := range address {
					fmt.Printf("adderss[%s%s%s], ", yellow, a.Addr, normal)
				}
				fmt.Printf("\n")
			}
		}

		return nil
	},
}
