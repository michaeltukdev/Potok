package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type VaultInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Config struct {
	APIURL   string      `json:"api_url"`
	Username string      `json:"username,omitempty"`
	Vaults   []VaultInfo `json:"vaults,omitempty"`
}

func configPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, ".potok")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", err
	}

	return filepath.Join(dir, "config.json"), nil
}

func Load() (*Config, error) {
	path, err := configPath()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return &Config{}, nil
	} else if err != nil {
		return nil, err
	}

	defer f.Close()
	var cfg Config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func Save(cfg *Config) error {
	path, err := configPath()
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(cfg)
}

func (cfg *Config) AddVault(vault VaultInfo) {
	for i, v := range cfg.Vaults {
		if v.Name == vault.Name {
			cfg.Vaults[i] = vault
			return
		}
	}

	cfg.Vaults = append(cfg.Vaults, vault)
}
