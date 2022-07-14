package config

import (
	"az-func-deploy/util"
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	__fileNameConfig string = "deploy.config.json"
)

func ReadConfigOrPanic() *DeployConfig {
	execDir, _ := os.Executable()
	configFile := filepath.Join(execDir, __fileNameConfig)
	if util.ExistsFile(configFile) {
		return readFromFile(configFile)
	}
	wd, _ := os.Getwd()
	configFile = filepath.Join(wd, __fileNameConfig)
	if util.ExistsFile(configFile) {
		return readFromFile(configFile)
	}
	WriteEmptyConfig(configFile)
	panic("Could not find config file, sample output: " + configFile)
}

func readFromFile(filePath string) *DeployConfig {
	config := DeployConfig{}
	bytes, err := os.ReadFile(filePath)
	util.PanicIfNotNil(err)
	err = json.Unmarshal(bytes, &config)
	util.PanicIfNotNil(err)
	config.ConfigJsonLocation = filePath
	return &config
}
