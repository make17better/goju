package config

import (
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

// Config represents the application configuration
type Config struct {
	Language string `yaml:"language"`
	Theme    string `yaml:"theme"`
	History  struct {
		Enabled bool `yaml:"enabled"`
		Limit   int  `yaml:"limit"`
	} `yaml:"history"`
	Practice struct {
		DefaultCount int      `yaml:"default_count"`
		Categories   []string `yaml:"categories"`
	} `yaml:"practice"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	cfg := &Config{
		Language: "en",
		Theme:    "default",
	}
	cfg.History.Enabled = true
	cfg.History.Limit = 100
	cfg.Practice.DefaultCount = 10
	cfg.Practice.Categories = []string{"seion", "dakuon", "handaku", "yoon"}
	return cfg
}

// GetConfigDir returns the configuration directory path
func GetConfigDir() (string, error) {
	var configDir string
	switch runtime.GOOS {
	case "windows":
		configDir = filepath.Join(os.Getenv("APPDATA"), "goju")
	case "darwin":
		configDir = filepath.Join(os.Getenv("HOME"), ".goju")
	default: // Linux and other Unix-like systems
		configDir = filepath.Join(os.Getenv("HOME"), ".goju")
	}

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}

	return configDir, nil
}

// LoadConfig loads the configuration from file
func LoadConfig() (*Config, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(configDir, "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		cfg := DefaultConfig()
		if err := SaveConfig(cfg); err != nil {
			return nil, err
		}
		return cfg, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	cfg := DefaultConfig()
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// SaveConfig saves the configuration to file
func SaveConfig(cfg *Config) error {
	configDir, err := GetConfigDir()
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	configPath := filepath.Join(configDir, "config.yaml")
	return os.WriteFile(configPath, data, 0644)
}

// GetHistoryPath returns the path to the history file
func GetHistoryPath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "history.yaml"), nil
}

// GetLogPath returns the path to the log file
func GetLogPath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "goju.log"), nil
}
