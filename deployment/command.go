package deployment

import (
	"az-func-deploy/logger"
	"az-func-deploy/util"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type PlatformDeployCommands interface {
	PrintBinaryVersions(logger *logger.Logger)
	DotNetBuild(projectDir string)
	ZipBuildOutput(outputDir string, projectDir string) string
	// https://docs.microsoft.com/en-us/azure/app-service/deploy-zip?tabs=cli
	AzureZipDeploy(resourceGroup string, funcName string, projectDir string, zipFile string)
	// https://docs.microsoft.com/en-us/cli/azure/functionapp/deployment/source?view=azure-cli-latest#az-functionapp-deployment-source-config-zip
	AzureFuncZipDeploy(resourceGroup string, funcName string, projectDir string, zipFile string)
	// https://docs.microsoft.com/en-us/azure/azure-functions/functions-run-local?tabs=v4%2Cwindows%2Ccsharp%2Cportal%2Cbash#project-file-deployment
	FuncDeployProject(funcName string, projectDir string)
}

func CommandStartAndWait(w io.Writer, dir string, name string, param ...string) {
	if util.GetIfShouldRun() {
		cmd := exec.Command(name, param...)
		util.SetKillFunc(func() { CommandKill(cmd.Process.Pid) })
		if dir != "" {
			cmd.Dir = dir
		}
		cmd.Stdout = w
		cmd.Stderr = w
		cmd.Start()
		cmd.Wait()
		util.SetKillFunc(nil)
	}
}

func CommandKill(pid int) {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command(getCmdExeLocation(), "/c", "TASKKILL", "/F", "/T", "/PID", fmt.Sprint(pid))
		cmd.Output()
	case "darwin":
		panic("Not implemented")
	case "linux":
		panic("Not implemented")
	default:
		panic("Current platform not supported")
	}
}

func NewCommandSet(writer io.Writer) PlatformDeployCommands {
	switch runtime.GOOS {
	case "windows":
		return &cmdCtxWindows{w: writer, cmdExe: getCmdExeLocation()}
	case "darwin":
		return &cmdCtxDarwin{w: writer}
	case "linux":
		return &cmdCtxLinux{w: writer}
	default:
		panic("Current platform not supported")
	}
}

func getCmdExeLocation() string {
	return filepath.Join(os.Getenv("SYSTEMROOT"), "System32", "cmd.exe")
}
