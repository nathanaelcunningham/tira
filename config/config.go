package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	User UserConfig `mapstructure:"user"`
    Server string `mapstructure:"server"`
}

type UserConfig struct {
	Username string `mapstructure:"username"`
	Token    string `mapstructure:"token"`
}

func CheckConfig() bool {
    home, _ := os.UserHomeDir()
    configPath := filepath.Join(home, ".config", "tira")
    configFile := filepath.Join(configPath, "tira_config.json")

    if _, err := os.Stat(configFile); os.IsNotExist(err) {
        return false
    }

    return true
}
func LoadConfig() {
    home, _ := os.UserHomeDir()
    configPath := filepath.Join(home, ".config", "tira")
    configFile := filepath.Join(configPath, "tira_config.json")

    viper.SetConfigName("tira_config")
    viper.SetConfigType("json")
    viper.AddConfigPath(configPath)

    if _, err := os.Stat(configFile); os.IsNotExist(err) {
        if err := os.MkdirAll(filepath.Dir(configFile), 0770); err != nil {
            log.Fatal("Error creating config file")
        }

        os.Create(configFile)
        viper.WriteConfig()
    }

    if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Error loading config:", err)
		}
	}
}

func GetConfig() (config Config) {
    err := viper.Unmarshal(&config)

    if err != nil {
        fmt.Println("Error parsing config", err)
    }
    return
}
