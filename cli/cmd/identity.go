package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(identityCmd)
}

var identityCmd = &cobra.Command{
	Use:   "identity",
	Short: "Identity commands",
}
