package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// 'BruteForce is a struct of yaml brute force'
type Config struct {
	BruteForce BruteForceConfig `yaml:"brute_force"`
}

type BruteForceConfig struct {
	Threshold int    `yaml:"threshold"`
	Window    string `yaml:"window"`
}

// reads a path like "configs/configs.yml" if any error found then sends it to Config{}
// return that config and error
func Load(path string) (Config, error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return Config{}, err
	}

	var config Config

	err = yaml.Unmarshal(data, &config)

	if err != nil {
		return Config{}, err
	}

	return config, nil
}
