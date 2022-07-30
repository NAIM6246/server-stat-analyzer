package configs

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type AppConfig struct {
	ListenPort int `json:"listenPort"`
}

const path = "/config.json"

var appConfig *AppConfig

func loadConfig() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filePath := fmt.Sprintf("%s%s", workingDir, path)
	stream, err := os.Open(filePath)
	if err != nil {

		panic(err)
	}
	parseErr := json.NewDecoder(stream).Decode(&appConfig)
	if parseErr != nil {
		panic(parseErr)
	}
}

//LoadConfig from json file
func init() {
	var loadOnce sync.Once
	loadOnce.Do(loadConfig)
}

func GetAppConfig() *AppConfig {
	return appConfig
}
