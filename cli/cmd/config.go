package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	ServiceURL *struct {
		Identity      string `yaml:"identity"`
		Account       string `yaml:"account"`
		Compute       string `yaml:"compute"`
		BlockStorage  string `yaml:"block_storage"`
		Image         string `yaml:"image"`
		Network       string `yaml:"network"`
		ObjectStorage string `yaml:"object_storage"`
		Database      string `yaml:"database"`
		DNS           string `yaml:"dns"`
		Mail          string `yaml:"mail"`
	} `yaml:"service_url"`
	User *struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		TenantId string `yaml:"tenant_id"`
		Token    string `yaml:"token"`
	} `yaml:"user"`
}

func loadConfig(path string) *config {
	var c *config

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
