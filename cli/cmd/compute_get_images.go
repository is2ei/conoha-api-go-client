package cmd

import (
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
		if client.Token == "" {
			access, _, err := client.IdentityToken()
			if err != nil {
				fmt.Println(err)
				return err
			} else {
				client.Token = access.Token.Id
			}
		}
		images, _, err := client.ComputeImages()
		if err != nil {
			fmt.Println(err)
			return err
		}

		if verbose {
			pp.Println(images)
		} else {
			for _, image := range images {
				fmt.Printf("name[%s%s%s], id[%s%s%s]\n", green, image.Name, normal, yellow, image.Id, normal)
			}
		}

		return nil
	},
}
