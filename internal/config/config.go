package config

import (
	"EquityEye/internal/logs"
	"EquityEye/types"
	"encoding/json"
	"errors"
	"os"
)

type EnvConfig struct {
	Environment bool // set this to true if running in a production env
	ConfigPath  string
}

type Config struct {
	Environment bool
	Providers   []types.ProviderConfiguration `yaml:"providers" json:"providers"`
	Cache       struct {
		Url string `yaml:"url" json:"url"`
	} `json:"cache"`
}

const (
	MARKETSTACK = "MARKETSTACK"
	BINANCE     = "BINANCE"
)

func LoadConfig() (*Config, error) {
	// Load ENV config
	config, err := loadEnvConfig()
	if err != nil {
		return nil, err
	}

	return loadYamlConfigFromFile(config)

}

func loadEnvConfig() (*EnvConfig, error) {
	config := &EnvConfig{}

	if env := os.Getenv("ENVIRONMENT"); env == "dev" {
		logs.Info("Using development environment")
		config.Environment = false
	} else {
		if env == "" {
			logs.Warn("Environment variable ENVIRONMENT NOT set, defaulting to prod")
		} else {
			logs.Info("Using production environment")
		}
		config.Environment = true
	}

	if env := os.Getenv("CONFIG_LOCATION"); env == "" {
		logs.Warn("Couldn't not find CONFIG_LOCATION environment variable, trying to locate it...")
		return nil, errors.New("couldn't load config from file")
	} else {
		logs.Info("Trying to load config from %s", env)
		config.ConfigPath = env
	}

	return config, nil
}

func loadYamlConfigFromFile(cfg *EnvConfig) (*Config, error) {
	data, err := os.ReadFile(cfg.ConfigPath)
	if err != nil {
		return nil, err
	}

	config := &Config{
		Environment: cfg.Environment,
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, errors.New("couldn't unmarshall config")
	}

	err = validateConfig(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func validateConfig(config *Config) error {
	if len(config.Providers) <= 0 {
		logs.Info("No providers are configured")
	}

	if config.Cache.Url == "" {
		return errors.New("no cache url provided (no cache setup will be supported in the future)")
	}

	names := make(map[string]bool)
	for _, provider := range config.Providers {
		if names[provider.Name] {
			return errors.New("provider name is duplicated")
		}
		names[provider.Name] = true

		if provider.Name == "" {
			return errors.New("provider name is required")
		}
		if provider.Key == "" {
			return errors.New("provider api key is required")
		}
		if provider.Limit <= 0 {
			logs.Info("Provider %s api limit is set to 0", provider.Name)
		}
		if provider.Limit > 0 && provider.LimitTimeframe <= 0 {
			return errors.New("provider api limit is set but no timeframe was given")
		}
	}

	logs.Info("Found %d available providers", len(config.Providers))

	return nil
}
