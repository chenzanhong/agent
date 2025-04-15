package configs

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Url           string `yaml:"url"`
	Port          string `yaml:"port"`
	Second        int    `yaml:"second"`
	FrpPublicIp   string `yaml:"frp_public_ip"`   // frp
	FrpRemotePort string `yaml:"frp_remote_port"` // frp
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config file: %v", err)
	}
	return &config, nil
}
