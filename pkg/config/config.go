package config

import (
	"github.com/1r0npipe/url-requestor/pkg/errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Server struct {
		Address  string  `yaml:"address"`
		Port     string  `yaml:"port"`
		Timeout  int     `yaml:"timeout"`
		LogLevel *string `yaml:"logLevel"`
	} `yaml:"server"`
	URLs    []string `yaml:"urls"`
	Workers int
}

func ReadConfigFile(configPath string) (*Config, error) {
	config := &Config{}
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, server_errors.ErrFileRead
	}
	if err := yaml.Unmarshal(file, config); err != nil {
		return nil, server_errors.ErrDecodeYAML
	}
	config.Workers = len(config.URLs)
	return config, nil
}
