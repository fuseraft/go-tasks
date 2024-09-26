package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

// Config structure for reading routes from the config file
type Config struct {
	Routes []RouteConfig `yaml:"routes"`
}

// RouteConfig defines a route in the config file
type RouteConfig struct {
	Path    string `yaml:"path"`
	Method  string `yaml:"method"`
	Handler string `yaml:"handler"`
}

// ReadConfig reads the YAML configuration file and returns a Config struct
func ReadConfig(configPath string) (Config, error) {
	var config Config

	file, err := os.Open(configPath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
