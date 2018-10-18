package cmd

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/is2ei/conoha-api-go-client/conoha"
	"github.com/spf13/cobra"
)

const (
	red    = "\x1b[31m"
	green  = "\x1b[32m"
	yellow = "\x1b[33m"
	normal = "\x1b[0m"
)

var (
	client     *conoha.Conoha
	configFile = getHomeDir() + "/.conoha-api-go-client.yml"
	verbose    bool
)

var rootCmd = &cobra.Command{
	Use:   "conoha",
	Short: "Yet another ConoHa API client",
	Long: `Yet another ConoHa API client
			built by is2ei`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config := loadConfig(configFile)
		client = conoha.NewConoha(
			config.IdentityServiceUrl,
			config.AccountServiceUrl,
			config.ComputeServiceUrl,
			config.BlockStorageServiceUrl,
			config.ImageServiceUrl,
			config.NetworkServiceUrl,
			config.ObjectStorageServiceUrl,
			config.DatabaseServiceUrl,
			config.DnsServiceUrl,
			config.MailServiceUrl,
			config.Username,
			config.Password,
			config.TenantId,
			config.Token,
		)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Yet another ConoHa API client")
	},
}

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
