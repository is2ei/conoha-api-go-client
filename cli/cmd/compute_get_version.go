package cmd

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	computeCmd.AddCommand(getComputeVersionCmd)
}

var getComputeVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get API Version",
	RunE: func(cmd *cobra.Command, args []string) error {
		if client.Token == "" {
			access, _, err := client.IdentityToken()
			if err != nil {
				fmt.Println(err)
				return err
			}
			client.Token = access.Token.Id
		}

		version, _, err := client.ComputeVersion()
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(version)

		return nil
	},
}
