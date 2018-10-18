package cmd

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	getComputeFlavorsCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	computeCmd.AddCommand(getComputeFlavorsCmd)
}

var getComputeFlavorsCmd = &cobra.Command{
	Use:   "flavors",
	Short: "Get flavors",
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
		flavors, _, err := client.ComputeFlavors()
		if err != nil {
			fmt.Println(err)
			return err
		}

		if verbose {
			pp.Println(flavors)
		} else {
			for _, flavor := range flavors {
				fmt.Printf("name[\x1b[32m%s\x1b[0m],id[\x1b[33m%s\x1b[0m]\n", flavor.Name, flavor.Id)
			}
		}

		return nil
	},
}
