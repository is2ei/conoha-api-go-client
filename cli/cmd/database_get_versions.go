package cmd

import (
	"context"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	databaseCmd.AddCommand(getDatabaseVersionsCmd)
}

var getDatabaseVersionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "Get API versions",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		version, _, err := client.DatabaseVersions(ctx)
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(version)

		return nil
	},
}
