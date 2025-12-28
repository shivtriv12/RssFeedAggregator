package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Db_Url            string `json:"db_url"`
	Current_User_Name string `json:"current_user_name"`
}

func Read() (*Config, error) {
	homeDirectoryPath, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("Error getting home directory %w", err)
	}
	fileContent, err := os.ReadFile(homeDirectoryPath + "/.gatorconfig.json")
	if err != nil {
		return nil, fmt.Errorf("Error reading config file %w", err)
	}
	config := &Config{}
	err = json.Unmarshal(fileContent, config)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling config file content %w", err)
	}
	return config, nil
}

func SetUser(C *Config) error {
	homeDirectoryPath, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Error getting home directory %w", err)
	}
	newConfig, err := json.Marshal(*C)
	if err != nil {
		return fmt.Errorf("Error in marshaling gatorconfig %w", err)
	}
	err = os.WriteFile(homeDirectoryPath+"/.gatorconfig.json", newConfig, 0644)
	if err != nil {
		return fmt.Errorf("Error in writing gatorconfig %w", err)
	}
	return nil
}
