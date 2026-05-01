package config

import (
	"aero/internal/logger"
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	Proxy     ProxyConfig      `yaml:"proxy"`
	Upstreams []UpstreamConfig `yaml:"upstreams"`
}

type ProxyConfig struct {
	Port                string `yaml:"port"`
	HealthCheckInterval int    `yaml:"healthcheck_interval"`
}

type UpstreamConfig struct {
	Url string `yaml:"url"`
}

func Load(path string, verbose bool) Config {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		logger.ErrorLogger(err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		logger.ErrorLogger(err)
	}

	return config
}
