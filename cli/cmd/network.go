package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(networkCmd)
}

var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Network commands",
}
