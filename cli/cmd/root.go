package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "conoha",
	Short: "Yet another ConoHa API client",
	Long: `Yet another ConoHa API client
			built by is2ei`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
