package cmd

import (
	"context"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	identityCmd.AddCommand(getVersionCmd)
}

var getVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get API Version",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		version, _, err := client.IdentityAPIVersion(ctx)
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(version)

		return nil
	},
}
