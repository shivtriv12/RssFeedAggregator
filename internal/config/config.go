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
		fmt.Println("Unable to find homedirectory path")
		return nil, err
	}
	fileContent, err := os.ReadFile(homeDirectoryPath + "/.gatorconfig.json")
	if err != nil {
		fmt.Println("Unable to read file")
		return nil, err
	}
	config := &Config{}
	err = json.Unmarshal(fileContent, config)
	if err != nil {
		fmt.Println("Unable to unmarshal")
		return nil, err
	}
	return config, nil
}

func SetUser(C *Config) error {
	homeDirectoryPath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Unable to find homedirectory path")
		return err
	}
	newConfig, err := json.Marshal(*C)
	if err != nil {
		fmt.Println("Error in marshaling Config")
		return err
	}
	err1 := os.WriteFile(homeDirectoryPath+"/.gatorconfig.json", newConfig, 0644)
	if err1 != nil {
		fmt.Println("Error in writing file")
		return err1
	}
	return nil
}
