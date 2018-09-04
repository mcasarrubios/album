package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/mcasarrubios/album/common"
)

type awsConfig struct {
	Endpoint string `json:"endPoint"`
	Region   string `json:"region"`
}

type dbConfig struct {
	PhotoTable string `json:"photoTable"`
	AlbumTable string `json:"albumTable"`
}

// Config data for the environment
type Config struct {
	APIURL string `json:"apiUrl"`
	DB     dbConfig
	AWS    awsConfig
}

// GetConfig gets the configuration
func GetConfig() *Config {
	env := os.Getenv("UP_STAGE")
	if env == "" {
		env = "development"
	}
	if env == "development" && os.Getenv("UP_TEST") == "true" {
		env = "test"
	}
	return getConfig(env)
}

func getConfig(env string) *Config {
	_, filename, _, _ := runtime.Caller(1)
	filePath := path.Join(path.Dir(filename), env+".json")
	jsonBytes, err := common.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading config file")
		os.Exit(1)
	}
	config := &Config{}
	err = json.Unmarshal(jsonBytes, config)
	if err != nil {
		fmt.Println("Error reading config file")
		os.Exit(1)
	}
	return config
}
