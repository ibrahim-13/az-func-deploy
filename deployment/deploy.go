package deployment

import (
	"az-func-deploy/config"
	"io"
)

func DeployFunctions(conf *config.DeployConfig, writer io.Writer) {
	writeHighlightln(writer, "Starting Deployment")
	for _, funcInfo := range conf.Sets[conf.CurrentSet].FuncInfos {
		writeBlackYellowln(writer, formatFuncMsg(funcInfo.FuncName, "Deploying Function"))
		if !funcInfo.ShouldRun {
			writeRedln(writer, formatFuncMsg(funcInfo.FuncName, "Skipped"))
			continue
		}
		CommandStartAndWait(writer, "cmd.exe", "/c", "dir")
		writeBlackYellowln(writer, formatFuncMsg(funcInfo.FuncName, "End"))
	}
	writeHighlightln(writer, "Finised Deployment")
}

func formatFuncMsg(name string, msg string) string {
	return "[" + name + "] " + msg
}
