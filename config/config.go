package config

import (
	"github.com/hjson/hjson-go/v4"
	"os"
)

const (
	EnvConfig = "GO_CONFIG"
)

type Configuration struct {
	Connections map[string]string `json:"connections"`
}

var currentConfig *Configuration

func isConfigFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// LoadConfig Find and load configuration file.
// File location can be provided via GO_CONFIG environment variable and must be JSON file.
// Default file is 'settings-example.json' and located in current working directory.
func LoadConfig() *Configuration {
	if currentConfig != nil {
		return currentConfig
	}

	return ReloadConfig()
}

// GetConfiguration Get current Configuration instance or load if it is not initialized yet
func GetConfiguration() *Configuration {
	if currentConfig == nil {
		LoadConfig()
	}

	return currentConfig
}

func ReloadConfig() *Configuration {

	cfgFile := os.Getenv(EnvConfig)

	if cfgFile == "" {
		cfgFile = "config.json"
	}

	if !isConfigFileExists(cfgFile) {
		panic("[ERROR] Could not file configuration file " + cfgFile + " \r\n Please provide " + EnvConfig + " environment variable")
	}

	buffer, err := os.ReadFile(cfgFile)

	if err != nil {
		panic("[ERROR] Error when read configuration file " + cfgFile + "\n" + err.Error())
	}

	currentConfig = &Configuration{}

	err = hjson.Unmarshal(buffer, currentConfig)

	if err != nil {
		panic("[ERROR] Error when parsing json configuration " + cfgFile + "\n" + err.Error())
	}

	return currentConfig
}
