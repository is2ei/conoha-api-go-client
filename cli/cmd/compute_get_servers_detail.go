package cmd

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	getComputeServersDetailCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	computeCmd.AddCommand(getComputeServersDetailCmd)
	rootCmd.AddCommand(getComputeServersDetailCmd) // alias
}

var getComputeServersDetailCmd = &cobra.Command{
	Use:   "servers-detail",
	Short: "Get servers detail",
	RunE: func(cmd *cobra.Command, args []string) error {
		if client.Token == "" {
			access, _, err := client.IdentityToken()
			if err != nil {
				fmt.Println(err)
				return err
			}
			client.Token = access.Token.ID
		}

		servers, _, err := client.ComputeServersDetail()
		if err != nil {
			fmt.Println(err)
			return err
		}

		if verbose {
			pp.Println(servers)
		} else {
			for _, server := range servers {
				fmt.Printf("name[%s%s%s], id[%s%s%s], status[%s%s%s], ",
					green, server.Name, normal, yellow, server.ID, normal, green, server.Status, normal)
				for _, address := range server.Addresses {
					for _, a := range address {
						fmt.Printf("address[%s%s%s], ", yellow, a.Addr, normal)
					}
					fmt.Printf("\n")
				}
			}
		}

		return nil
	},
}
