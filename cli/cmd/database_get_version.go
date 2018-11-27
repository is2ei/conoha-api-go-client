package cmd

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	databaseCmd.AddCommand(getDatabaseVersionCmd)
}

var getDatabaseVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get API version",
	RunE: func(cmd *cobra.Command, args []string) error {
		if client.Token == "" {
			access, _, err := client.IdentityToken()
			if err != nil {
				fmt.Println(err)
				return err
			}
			client.Token = access.Token.ID
		}
		version, _, err := client.DatabaseVersion()
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(version)

		return nil
	},
}
