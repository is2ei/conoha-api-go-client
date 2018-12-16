package cmd

import (
	"context"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	getComputeServersCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	computeCmd.AddCommand(getComputeServersCmd)
	rootCmd.AddCommand(getComputeServersCmd) // alias
}

var getComputeServersCmd = &cobra.Command{
	Use:   "servers",
	Short: "Get servers",
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

		servers, _, err := client.ComputeServers(ctx)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if verbose {
			pp.Println(servers)
		} else {
			for _, server := range servers {
				fmt.Printf("name[%s%s%s], id[%s%s%s]\n", green, server.Name, normal, yellow, server.ID, normal)
			}
		}

		return nil
	},
}
