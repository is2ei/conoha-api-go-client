package cmd

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

var (
	imageRef          string
	flavorRef         string
	adminPass         string
	keyName           string
	securityGroupName string
)

func init() {
	addComputeServerCmd.Flags().StringVarP(&imageRef, "image", "i", "", "Image UUID")
	addComputeServerCmd.Flags().StringVarP(&flavorRef, "flavor", "f", "", "VM Plan(Flavor) UUID")
	addComputeServerCmd.Flags().StringVarP(&adminPass, "adminPass", "p", "", "Root password")
	addComputeServerCmd.Flags().StringVarP(&keyName, "keyName", "k", "", "SSH Key name")
	addComputeServerCmd.Flags().StringVarP(&securityGroupName, "securityGroupName", "s", "", "Security Group Name")
	addComputeServerCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	serverCmd.AddCommand(addComputeServerCmd)
}

var addComputeServerCmd = &cobra.Command{
	Use:   "add",
	Short: "Add server",
	RunE: func(cmd *cobra.Command, args []string) error {
		if client.Token == "" {
			access, _, err := client.IdentityToken()
			if err != nil {
				fmt.Println(err)
				return err
			}
			client.Token = access.Token.Id
		}
		server, _, err := client.AddComputeServer(
			imageRef,
			flavorRef,
			adminPass,
			keyName,
			securityGroupName,
		)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if verbose {
			pp.Println(server)
		} else {
			fmt.Printf("id[%s%s%s], password[%s%s%s]", green, server.Id, normal, yellow, server.AdminPass, normal)
			fmt.Printf("\n")
		}

		return nil
	},
}
