package cmd

import (
	"context"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func init() {
	mailCmd.AddCommand(getMailVersionCmd)
}

var getMailVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get API version",
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
		version, _, err := client.MailVersion(ctx)
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(version)

		return nil
	},
}
