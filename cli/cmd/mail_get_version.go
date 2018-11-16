package cmd

import (
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
		if client.Token == "" {
			access, _, err := client.IdentityToken()
			if err != nil {
				fmt.Println(err)
				return err
			}
			client.Token = access.Token.Id
		}
		version, _, err := client.MailVersion()
		if err != nil {
			fmt.Println(err)
			return err
		}

		pp.Println(version)

		return nil
	},
}
