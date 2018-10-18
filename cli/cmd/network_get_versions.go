package cmd

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	networkCmd.AddCommand(getNetworkVersionsCmd)
}

var getNetworkVersionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "Get API ",
	RunE: func(cmd *cobra.Command, args []string) error {
		versions, _, err := client.NetworkVersions()
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(versions)

		return nil
	},
}
