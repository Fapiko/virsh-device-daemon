package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v1"
)

/*
Config holds the configuration data for running the app in daemon mode
*/
type Config struct {
	Name        string   `yaml:"name"`
	Devices     string   `yaml:"devices"`
	DeviceFiles []string `yaml:"-"`
}

func parseConfig(path string) (*Config, error) {
	configData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(configData, config)
	if err != nil {
		return nil, err
	}

	if config.Devices != "" {
		files, err := ioutil.ReadDir(config.Devices)
		if err != nil {
			return nil, err
		}

		for _, f := range files {
			if f.IsDir() == false {
				config.DeviceFiles = append(config.DeviceFiles, fmt.Sprintf("%s%s%s",
					config.Devices, string(os.PathSeparator), f.Name()))
			}
		}
	}

	return config, err
}
