package cmd

import (
	"context"
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
		ctx := context.Background()
		if client.Token == "" {
			access, _, err := client.IdentityToken(ctx)
			if err != nil {
				fmt.Println(err)
				return err
			}
			client.Token = access.Token.ID
		}
		flavors, _, err := client.ComputeFlavors(ctx)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if verbose {
			pp.Println(flavors)
		} else {
			for _, flavor := range flavors {
				fmt.Printf("name[\x1b[32m%s\x1b[0m],id[\x1b[33m%s\x1b[0m]\n", flavor.Name, flavor.ID)
			}
		}

		return nil
	},
}
