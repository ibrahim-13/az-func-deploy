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
	cmds := NewCommandSet(writer)
	logger.Highlightln("Starting Deployment")
	for _, funcInfo := range currentSet.FuncInfos {
		logger.SetScope(funcInfo.FuncName)
		logger.BlackYellowln("Deploying Function")
		if !funcInfo.ShouldRun {
			logger.Redln("Skipped")
			continue
		}
		if conf.Method == config.DeployMethodFunc {
			cmds.FuncDeployProject(funcInfo.FuncName, funcInfo.ProjectDir)
		} else {
			cmds.DotNetBuild(funcInfo.ProjectDir)
			baseConfigDir := filepath.Dir(conf.ConfigJsonLocation)
			outputFile := cmds.ZipBuildOutput(baseConfigDir, funcInfo.ProjectDir)
			if conf.Method == config.DeployMethodAzFunc {
				cmds.AzureFuncZipDeploy(currentSet.ResourceGroupName, funcInfo.FuncName, funcInfo.ProjectDir, outputFile)
			} else if conf.Method == config.DeployMethodAzZip {
				cmds.AzureFuncZipDeploy(currentSet.ResourceGroupName, funcInfo.FuncName, funcInfo.ProjectDir, outputFile)
			} else {
				panic("Invalid deployment methdo: " + conf.Method)
			}
			os.Remove(outputFile)
		}
		logger.BlackYellowln("End")
	}
	logger.Highlightln("Finised Deployment")
}
