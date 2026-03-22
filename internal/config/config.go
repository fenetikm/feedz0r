package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const configFileName = "config.yml"

type FetchConfig struct {
	Timeout     int `yaml:"timeout"`
	RefreshMins int `yaml:"refresh"`
	MaxFeeds    int `yaml:"maxfeeds"`
}

type Config struct {
	Debug bool        `yaml:"debug"`
	Fetch FetchConfig `yaml:"fetch"`
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
