package config

import (
	"az-func-deploy/util"
	"encoding/json"
	"os"
)

func WriteEmptyConfig(filePath string) {
	funcInfo := []DeployFuncInfo{{}}
	sets := []DeploySet{{FuncInfos: funcInfo}}
	config := DeployConfig{Sets: sets, Method: DeployMethodAzFunc}
	bytes, err := json.Marshal(config)
	util.PanicIfNotNil(err)
	err = os.WriteFile(filePath, bytes, 0777)
	util.PanicIfNotNil(err)
}

func (c *DeployConfig) WriteConfig() {
	bytes, err := json.MarshalIndent(c, "", "  ")
	util.PanicIfNotNil(err)
	err = os.WriteFile(c.ConfigJsonLocation, bytes, 0755)
	util.PanicIfNotNil(err)
}
