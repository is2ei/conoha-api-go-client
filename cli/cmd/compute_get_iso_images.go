package cmd

import (
	"context"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	getComputeIsoImagesCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	computeCmd.AddCommand(getComputeIsoImagesCmd)
}

var getComputeIsoImagesCmd = &cobra.Command{
	Use:   "iso-list",
	Short: "Get ISO images list",
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
		isoImages, _, err := client.IsoImages(ctx)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if verbose {
			pp.Println(isoImages)
		} else {
			for _, isoImage := range isoImages {
				fmt.Printf("name[%s]\n", isoImage.Name)
			}
		}

		return nil
	},
}
