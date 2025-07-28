package config

import "fmt"

func MustLoadWithAPIURL() (*Config, error) {
	cfg, err := Load()
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	if cfg.APIURL == "" {
		return nil, fmt.Errorf("API URL is not set. Please run 'potok set-api-url' first.")
	}
	return cfg, nil
}
