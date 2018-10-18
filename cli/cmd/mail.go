package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(mailCmd)
}

var mailCmd = &cobra.Command{
	Use:   "mail",
	Short: "Mail commands",
}
