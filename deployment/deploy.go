package deployment

import (
	"az-func-deploy/config"
	"az-func-deploy/logger"
	"fmt"
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
	totalFuncs := len(currentSet.FuncInfos)
	deployedFuncs := []string{}
	for i, funcInfo := range currentSet.FuncInfos {
		logger.SetScope(fmt.Sprintf("%d/%d | %s", i+1, totalFuncs, funcInfo.FuncName))
		logger.ScopedBlackYellowln("Deploying Function")
		logger.BlackYellowln(funcInfo.ProjectDir)
		if !funcInfo.ShouldRun {
			logger.ScopedBlackRedln("Skipped")
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
		deployedFuncs = append(deployedFuncs, funcInfo.FuncName)
		logger.ScopedBlackYellowln("End")
	}
	logger.Highlightln(fmt.Sprintf("Deployed %d of %d functions", len(deployedFuncs), totalFuncs))
	for _, name := range deployedFuncs {
		logger.Greenln(name)
	}
	logger.Highlightln("Finised Deployment")
}
