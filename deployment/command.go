package deployment

import (
	"io"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func CommandStartAndWait(w io.Writer, dir string, name string, param ...string) {
	cmd := exec.Command(name, param...)
	cmd.Dir = dir
	cmd.Stdout = w
	cmd.Stderr = w
	cmd.Start()
	cmd.Wait()
}

func CommandDotNetBuild(w io.Writer, projectDir string) {
	CommandStartAndWait(w, projectDir, "C:\\Windows\\System32\\cmd.exe", "/c", "dotnet", "build", "--configuration", "Release")
}

func CommandZipBuildOutput(w io.Writer, outputDir string, projectDir string) string {
	buildDir := filepath.Join(projectDir, "bin", "Release", "net6.0")
	zipFile := filepath.Join(outputDir, strconv.FormatInt(time.Now().Unix(), 10)+".zip")
	CommandStartAndWait(w, buildDir, "C:\\Windows\\System32\\cmd.exe", "/c", "powershell", "-command", "Compress-Archive -Path * -DestinationPath "+zipFile)
	return zipFile
}

// https://docs.microsoft.com/en-us/azure/app-service/deploy-zip?tabs=cli
func CommandAzureZipDeploy(w io.Writer, resourceGroup string, funcName string, projectDir string, zipFile string) {
	CommandStartAndWait(w, projectDir, "C:\\Windows\\System32\\cmd.exe", "/c", "az", "webapp", "deploy", "--resource-group", resourceGroup, "--name", funcName, "--src-path", zipFile)
}

// https://docs.microsoft.com/en-us/cli/azure/functionapp/deployment/source?view=azure-cli-latest#az-functionapp-deployment-source-config-zip
func CommandAzureFuncZipDeploy(w io.Writer, resourceGroup string, funcName string, projectDir string, zipFile string) {
	CommandStartAndWait(w, projectDir, "C:\\Windows\\System32\\cmd.exe", "/c", "az", "functionapp", "deployment", "source", "config-zip", "--resource-group", resourceGroup, "--name", funcName, "--src", zipFile)
}

// https://docs.microsoft.com/en-us/azure/azure-functions/functions-run-local?tabs=v4%2Cwindows%2Ccsharp%2Cportal%2Cbash#project-file-deployment
func CommandFuncDeployProject(w io.Writer, funcName string, projectDir string) {
	CommandStartAndWait(w, projectDir, "C:\\Windows\\System32\\cmd.exe", "/c", "func", "azure", "functionapp", "publish", funcName)
}
