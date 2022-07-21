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
	startTime := time.Now()
	logger := logger.NewLogger(writer)
	logger.SetColor(!disableColor)
	currentSet := conf.Sets[conf.CurrentSet]
	cmds := NewCommandSet(writer)
	cmds.PrintBinaryVersions(logger)
	logger.Highlightln("Deployment will start in 5 seconds...")
	time.Sleep(5 * time.Second)
	logger.Highlightln("DEPLOYMENT START")
	totalFuncs := len(currentSet.FuncInfos)
	deployedFuncs := []string{}
	skippedFuncs := []string{}
	cmdCount := 0
	cmdOkCount := 0
	handleCmdResult := func(ok bool) {
		cmdCount += 1
		if ok {
			cmdOkCount += 1
			logger.ScopedWhiteGreenln("Success")
		} else {
			logger.ScopedBlackRedln("Exit with error")
		}
	}
	for i, funcInfo := range currentSet.FuncInfos {
		logger.SetScope(fmt.Sprintf("%2d/%2d | %s", i+1, totalFuncs, funcInfo.FuncName))
		logger.ScopedBlackYellowln("Deploying Function")
		logger.BlackYellowln(funcInfo.ProjectDir)
		if !funcInfo.ShouldRun {
			logger.ScopedBlackRedln("Skipped")
			skippedFuncs = append(skippedFuncs, funcInfo.FuncName)
			continue
		}
		logger.ScopedWhiteBlueln("Deploy method: " + conf.Method)
		if conf.Method == config.DeployMethodFunc {
			ok := cmds.FuncDeployProject(funcInfo.FuncName, funcInfo.ProjectDir)
			handleCmdResult(ok)
		} else {
			logger.ScopedWhiteBlueln("Building project...")
			ok := cmds.DotNetBuild(funcInfo.ProjectDir)
			handleCmdResult(ok)
			baseConfigDir := filepath.Dir(conf.ConfigJsonLocation)
			logger.ScopedWhiteBlueln("Creating zip artifact...")
			outputFile, ok := cmds.ZipBuildOutput(baseConfigDir, funcInfo.ProjectDir)
			handleCmdResult(ok)
			logger.ScopedWhiteBlueln("Deploying artifact...")
			if conf.Method == config.DeployMethodAzFunc {
				ok := cmds.AzureFuncZipDeploy(currentSet.ResourceGroupName, funcInfo.FuncName, funcInfo.ProjectDir, outputFile)
				handleCmdResult(ok)
			} else if conf.Method == config.DeployMethodAzZip {
				ok := cmds.AzureZipDeploy(currentSet.ResourceGroupName, funcInfo.FuncName, funcInfo.ProjectDir, outputFile)
				handleCmdResult(ok)
			} else {
				panic("Invalid deployment methdo: " + conf.Method)
			}
			os.Remove(outputFile)
		}
		deployedFuncs = append(deployedFuncs, funcInfo.FuncName)
		logger.ScopedBlackYellowln("End")
	}
	logger.Highlightln(fmt.Sprintf("Deployed %d of %d functions", len(deployedFuncs), totalFuncs))
	for _, deployedFuncName := range deployedFuncs {
		logger.Greenln(deployedFuncName)
	}
	logger.Highlightln(fmt.Sprintf("Skipped %d of %d functions", len(skippedFuncs), totalFuncs))
	for _, skippedFuncName := range skippedFuncs {
		logger.Redln(skippedFuncName)
	}
	logger.Highlightln(fmt.Sprintf("Commands Succeed: %d/%d", cmdOkCount, cmdCount))
	logger.Highlightln(fmt.Sprintf("Commands Failed: %d/%d", cmdCount-cmdOkCount, cmdCount))
	logger.Highlightln(fmt.Sprintf("Elapsed Time: %s", time.Since(startTime)))
	logger.Highlightln("DEPLOYMENT END")
}
