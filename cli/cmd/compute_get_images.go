package cmd

import (
	"context"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	getComputeImagesCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	computeCmd.AddCommand(getComputeImagesCmd)
}

var getComputeImagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Get images",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		if client.Token == "" {
			access, _, err := client.IdentityToken(ctx)
			if err != nil {
				fmt.Println(err)
				return err
			}
			client.Token = access.Token.ID
		}
		images, _, err := client.ComputeImages(ctx)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if verbose {
			pp.Println(images)
		} else {
			for _, image := range images {
				fmt.Printf("name[%s%s%s], id[%s%s%s]\n", green, image.Name, normal, yellow, image.ID, normal)
			}
		}

		return nil
	},
}
