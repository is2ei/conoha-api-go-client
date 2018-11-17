package cmd

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	computeCmd.AddCommand(getComputeVersionsCmd)
}

var getComputeVersionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "Get API Versions",
	RunE: func(cmd *cobra.Command, args []string) error {
		if client.Token == "" {
			access, _, err := client.IdentityToken()
			if err != nil {
				fmt.Println(err)
				return err
			}
			client.Token = access.Token.Id
		}

		versions, _, err := client.ComputeVersions()
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(versions)

		return nil
	},
}
