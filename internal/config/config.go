package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	configFileName = ".gatorconfig.json"
	DefaultDbURL   = "postgres://example"
)

// Config represents the structure of the configuration file
type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// Read reads the JSON config file and returns a Config struct
// If the file doesn't exist, it creates a default one
func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	
	// Check if file exists, create default if it doesn't
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		defaultConfig := Config{
			DbURL: DefaultDbURL,
		}
		if writeErr := write(defaultConfig); writeErr != nil {
			return Config{}, writeErr
		}
		return defaultConfig, nil
	}
	
	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}
	
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}
	
	return cfg, nil
}

// SetUser sets the current user name and writes the config to disk
func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(*c)
}

// getConfigFilePath returns the full path to the config file
func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	
	return filepath.Join(homeDir, configFileName), nil
}

// write writes the config struct to the JSON file
func write(cfg Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(configPath, data, 0644)
}