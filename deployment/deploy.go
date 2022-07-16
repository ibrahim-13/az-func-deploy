package deployment

import (
	"az-func-deploy/config"
	"az-func-deploy/logger"
	"io"
	"os"
	"path/filepath"
)

func DeployFunctions(conf *config.DeployConfig, writer io.Writer) {
	logger := logger.NewLogger(writer)
	currentSet := conf.Sets[conf.CurrentSet]
	logger.Highlightln("Starting Deployment")
	for _, funcInfo := range currentSet.FuncInfos {
		logger.SetScope(funcInfo.FuncName)
		logger.BlackYellowln("Deploying Function")
		if !funcInfo.ShouldRun {
			logger.Redln("Skipped")
			continue
		}
		if conf.Method == config.DeployMethodFunc {
			CommandFuncDeployProject(writer, funcInfo.FuncName, funcInfo.ProjectDir)
		} else {
			CommandDotNetBuild(writer, funcInfo.ProjectDir)
			baseConfigDir := filepath.Dir(conf.ConfigJsonLocation)
			outputFile := CommandZipBuildOutput(writer, baseConfigDir, funcInfo.ProjectDir)
			if conf.Method == config.DeployMethodAzFunc {
				CommandAzureFuncZipDeploy(writer, currentSet.ResourceGroupName, funcInfo.FuncName, funcInfo.ProjectDir, outputFile)
			} else if conf.Method == config.DeployMethodAzZip {
				CommandAzureFuncZipDeploy(writer, currentSet.ResourceGroupName, funcInfo.FuncName, funcInfo.ProjectDir, outputFile)
			} else {
				panic("Invalid deployment methdo: " + conf.Method)
			}
			os.Remove(outputFile)
		}
		logger.BlackYellowln("End")
	}
	logger.Highlightln("Finised Deployment")
}
