package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(imageCmd)
}

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Image commands",
}
