package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`

	Server struct {
		Port    string `yaml:"port"`
		Timeout int    `yaml:"timeout"`
	} `yaml:"server"`
}

func LoadConfig() *Config {
	// Đường dẫn đến tệp tin database.yaml
	filePath := "/Users/admin/Documents/Developer/GoLang/lore_project/configs/config.yaml"

	// Read the config.yaml file
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	// Parse the configuration file
	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	return config
}
