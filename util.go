package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	Targets []string `yaml:"targets"`
}

func LoadTargets(filePath string) ([]string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read targets file: %v", err)
	}

	var config YamlConfig
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %v", err)
	}

	var cleanLinks []string
	for _, link := range config.Targets {
		trimmed := strings.TrimSpace(link)
		if trimmed != "" {
			cleanLinks = append(cleanLinks, trimmed)
		}
	}
	return cleanLinks, nil
}

func EnsureDir(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			return fmt.Errorf("could not create directory %s: %v", dirName, err)
		}
	}
	return nil
}

func SaveDataToFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}

func SetupLogger(logPath string) (*os.File, error) {
	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}
	log.SetOutput(file)
	return file, nil
}
