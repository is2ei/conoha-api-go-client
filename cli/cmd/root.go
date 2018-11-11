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
	normal = "\x1b[0m"
	red    = "\x1b[31m"
	green  = "\x1b[32m"
	yellow = "\x1b[33m"
	blue   = "\x1b[34m"
)

var (
	client     *conoha.Conoha
	configFile = getHomeDir() + "/.conoha-api-go-client.conf"
	verbose    bool
)

var rootCmd = &cobra.Command{
	Use:   "conoha",
	Short: "Yet another ConoHa API client",
	Long:  `Yet another ConoHa API client built by is2ei`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config := loadConfig(configFile)
		client = conoha.NewConoha(
			config.ServiceUrl.Identity,
			config.ServiceUrl.Account,
			config.ServiceUrl.Compute,
			config.ServiceUrl.BlockStorage,
			config.ServiceUrl.Image,
			config.ServiceUrl.Network,
			config.ServiceUrl.ObjectStorage,
			config.ServiceUrl.Database,
			config.ServiceUrl.DNS,
			config.ServiceUrl.Mail,
			config.User.Username,
			config.User.Password,
			config.User.TenantId,
			config.User.Token,
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
