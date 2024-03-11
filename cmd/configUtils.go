/*
Copyright © 2024 Iván Recasens Lahoz <irecasensl@gmail.com>
*/

package cmd

import (
	"encoding/json"
	"os"
	"strings"
)

// Provides utilities to access configuration file

var configFileName string = "\\.todoConfig.json"

type TodoConfig struct {
	VaultDirectory string
	Filename       string
}

func (c *TodoConfig) prettyConfig() string {
	var result strings.Builder

	result.WriteString("FILE CONFIGURATION\n")
	result.WriteString("===================\n")
	result.WriteString("Directory: " + c.VaultDirectory + "\n")
	result.WriteString("Filename: " + c.Filename + "\n")

	return result.String()
}

func GetConfigFileRaw() string {
	return string(getConfigData())
}

func GetConfigFilePretty() string {
	return getConfigInfo().prettyConfig()
}

func GetVaultDirectory() string {
	return getConfigInfo().VaultDirectory
}

func GetFilename() string {
	return getConfigInfo().Filename
}

func getConfigData() []byte {
	_, err := os.Stat(getHomeDirectory() + configFileName)
	if os.IsNotExist(err) {
		panic("Configuration needs to be set")
	}

	fileData, err := os.ReadFile(getHomeDirectory() + configFileName)
	if err != nil {
		panic("Error reading configuration file")
	}

	return fileData
}

func getConfigInfo() *TodoConfig {
	var cf TodoConfig
	
	err := json.Unmarshal(getConfigData(), &cf)
	if err != nil {
		panic("Error parsing configuration file, format may be corrupted")
	}

	return &cf
}

func getHomeDirectory() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("Failed to obtain user home directory")
	}

	return home
}
