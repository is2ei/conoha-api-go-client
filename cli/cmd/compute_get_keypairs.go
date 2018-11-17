package cmd

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	getComputeKeypairsCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	computeCmd.AddCommand(getComputeKeypairsCmd)
}

var getComputeKeypairsCmd = &cobra.Command{
	Use:   "keypairs",
	Short: "Get keypairs",
	RunE: func(cmd *cobra.Command, args []string) error {
		if client.Token == "" {
			access, _, err := client.IdentityToken()
			if err != nil {
				fmt.Println(err)
				return err
			}
			client.Token = access.Token.Id
		}

		keypairs, _, err := client.ComputeKeypairs()
		if err != nil {
			fmt.Println(err)
			return err
		}

		if verbose {
			pp.Println(keypairs)
		} else {
			for _, keypair := range keypairs {
				fmt.Printf("name[\x1b[32m%s\x1b[0m]", keypair.Value.Name)
			}
		}

		return nil
	},
}
