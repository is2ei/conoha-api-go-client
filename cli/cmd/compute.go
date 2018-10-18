package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(computeCmd)
}

var computeCmd = &cobra.Command{
	Use:   "compute",
	Short: "Compute commands",
}
