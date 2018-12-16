package cmd

import (
	"context"
	"fmt"

	"github.com/is2ei/conoha-api-go-client/conoha"
	"github.com/spf13/cobra"
)

func init() {
	getTokenCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	identityCmd.AddCommand(getTokenCmd)
	rootCmd.AddCommand(getTokenCmd) // alias
}

var getTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Get token",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		access, _, err := client.IdentityToken(ctx)
		if err != nil {
			switch err.(type) {
			case *conoha.ErrUnauthorized:
				fmt.Printf("%s%s%s\n", red, err, normal)
				return nil
			default:
				return err
			}
		}

		if verbose {
			fmt.Printf("%stoken%s[%s%s%s], ", yellow, normal, green, access.Token.ID, normal)
			fmt.Printf("%sexpires%s[%s%s%s]", yellow, normal, green, access.Token.Expires, normal)
			fmt.Printf("\n")
		} else {
			fmt.Print(access.Token.ID)
		}

		return nil
	},
}
