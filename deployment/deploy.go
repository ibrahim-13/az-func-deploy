package deployment

import (
	"az-func-deploy/config"
	"az-func-deploy/logger"
	"io"
	"os"
	"path/filepath"
	"time"
)

func DeployFunctions(conf *config.DeployConfig, writer io.Writer, disableColor bool) {
	logger := logger.NewLogger(writer)
	logger.SetColor(!disableColor)
	currentSet := conf.Sets[conf.CurrentSet]
	cmds := NewCommandSet(writer)
	cmds.PrintBinaryVersions(logger)
	logger.Highlightln("Deployment will start in 5 seconds...")
	time.Sleep(5 * time.Second)
	logger.Highlightln("Starting Deployment")
	for _, funcInfo := range currentSet.FuncInfos {
		logger.SetScope(funcInfo.FuncName)
		logger.BlackYellowln("Deploying Function: " + funcInfo.FuncName)
		if !funcInfo.ShouldRun {
			logger.BlackRedln("Skipped")
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
