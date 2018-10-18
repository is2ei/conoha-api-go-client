package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	IdentityServiceUrl      string `yaml:"identity_service_url"`
	AccountServiceUrl       string `yaml:"account_service_url"`
	ComputeServiceUrl       string `yaml:"compute_service_url"`
	BlockStorageServiceUrl  string `yaml:"block_storage_service_url"`
	ImageServiceUrl         string `yaml:"image_service_url"`
	NetworkServiceUrl       string `yaml:"network_service_url"`
	ObjectStorageServiceUrl string `yaml:"object_storage_service_url"`
	DatabaseServiceUrl      string `yaml:"database_service_url"`
	DnsServiceUrl           string `yaml:"dns_service_url"`
	MailServiceUrl          string `yaml:"mail_service_url"`
	Username                string `yaml:"username"`
	Password                string `yaml:"password"`
	TenantId                string `yaml:"tenant_id"`
	Token                   string `yaml:"token"`
}

func loadConfig(path string) *Config {
	var c *Config

	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(file, &c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return c
}
