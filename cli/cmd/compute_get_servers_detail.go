package cmd

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	getComputeServersDetailCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	computeCmd.AddCommand(getComputeServersDetailCmd)
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
			} else {
				client.Token = access.Token.Id
			}
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
				fmt.Printf("name[\x1b[32m%s\x1b[0m],id[\x1b[33m%s\x1b[0m],status[\x1b[34m%s\x1b[0m]", server.Name, server.Id, server.Status)
			}
		}

		return nil
	},
}
