package deployment

import (
	"az-func-deploy/config"
	"io"
	"os"
	"path/filepath"
)

func DeployFunctions(conf *config.DeployConfig, writer io.Writer) {
	writeHighlightln(writer, "Starting Deployment")
	for _, funcInfo := range conf.Sets[conf.CurrentSet].FuncInfos {
		writeBlackYellowln(writer, formatFuncMsg(funcInfo.FuncName, "Deploying Function"))
		if !funcInfo.ShouldRun {
			writeRedln(writer, formatFuncMsg(funcInfo.FuncName, "Skipped"))
			continue
		}
		if conf.Method == config.DeployMethodFunc {
			CommandFuncDeployProject(writer, funcInfo.FuncName, funcInfo.ProjectDir)
		} else {
			CommandDotNetBuild(writer, funcInfo.ProjectDir)
			baseConfigDir := filepath.Dir(conf.ConfigJsonLocation)
			outputFile := CommandZipBuildOutput(writer, baseConfigDir, funcInfo.ProjectDir)
			if conf.Method == config.DeployMethodAzFunc {
				CommandAzureFuncZipDeploy(writer, conf.Sets[conf.CurrentSet].ResourceGroupName, funcInfo.FuncName, funcInfo.ProjectDir, outputFile)
			} else if conf.Method == config.DeployMethodAzZip {
				CommandAzureFuncZipDeploy(writer, conf.Sets[conf.CurrentSet].ResourceGroupName, funcInfo.FuncName, funcInfo.ProjectDir, outputFile)
			} else {
				panic("Invalid deployment methdo: " + conf.Method)
			}
			os.Remove(outputFile)
		}
		writeBlackYellowln(writer, formatFuncMsg(funcInfo.FuncName, "End"))
	}
	writeHighlightln(writer, "Finised Deployment")
}

func formatFuncMsg(name string, msg string) string {
	return "[" + name + "] " + msg
}
