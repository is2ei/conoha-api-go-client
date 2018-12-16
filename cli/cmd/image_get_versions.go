package cmd

import (
	"context"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	imageCmd.AddCommand(getImageVersionsCmd)
}

var getImageVersionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "Get API versions",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		versions, _, err := client.ImageVersions(ctx)
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(versions)

		return nil
	},
}
