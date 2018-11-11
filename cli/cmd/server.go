package cmd

import "github.com/spf13/cobra"

func init() {
	computeCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(serverCmd) // alias
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server commands",
}
