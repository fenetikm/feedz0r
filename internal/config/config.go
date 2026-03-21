package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const configFileName = "config.yml"

type Config struct {
	Debug bool `yaml:"debug"`
}

func Load() (Config, error) {
	data, err := os.ReadFile(configFileName)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
