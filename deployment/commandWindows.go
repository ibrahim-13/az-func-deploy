package deployment

import (
	"az-func-deploy/logger"
	"io"
	"path/filepath"
	"strconv"
	"time"
)

type cmdCtxWindows struct {
	w io.Writer
}

func (ctx *cmdCtxWindows) PrintBinaryVersions(logger *logger.Logger) {
	logger.WhiteBlueln("dotnet --version")
	CommandStartAndWait(ctx.w, "", "C:\\Windows\\System32\\cmd.exe", "/c", "dotnet", "--version")
	logger.WhiteBlueln("az version --output tsv")
	CommandStartAndWait(ctx.w, "", "C:\\Windows\\System32\\cmd.exe", "/c", "az", "version", "--output", "tsv")
	logger.WhiteBlueln("func --version")
	CommandStartAndWait(ctx.w, "", "C:\\Windows\\System32\\cmd.exe", "/c", "dotnet", "--version")
}

func (ctx *cmdCtxWindows) DotNetBuild(projectDir string) {
	CommandStartAndWait(ctx.w, projectDir, "C:\\Windows\\System32\\cmd.exe", "/c", "dotnet", "build", "--configuration", "Release")
}

func (ctx *cmdCtxWindows) ZipBuildOutput(outputDir string, projectDir string) string {
	buildDir := filepath.Join(projectDir, "bin", "Release", "net6.0")
	zipFile := filepath.Join(outputDir, strconv.FormatInt(time.Now().Unix(), 10)+".zip")
	CommandStartAndWait(ctx.w, buildDir, "C:\\Windows\\System32\\cmd.exe", "/c", "powershell", "-command", "Compress-Archive -Path * -DestinationPath "+zipFile)
	return zipFile
}

func (ctx *cmdCtxWindows) AzureZipDeploy(resourceGroup string, funcName string, projectDir string, zipFile string) {
	CommandStartAndWait(ctx.w, projectDir, "C:\\Windows\\System32\\cmd.exe", "/c", "az", "webapp", "deploy", "--resource-group", resourceGroup, "--name", funcName, "--src-path", zipFile)
}

func (ctx *cmdCtxWindows) AzureFuncZipDeploy(resourceGroup string, funcName string, projectDir string, zipFile string) {
	CommandStartAndWait(ctx.w, projectDir, "C:\\Windows\\System32\\cmd.exe", "/c", "az", "functionapp", "deployment", "source", "config-zip", "--resource-group", resourceGroup, "--name", funcName, "--src", zipFile)
}

func (ctx *cmdCtxWindows) FuncDeployProject(funcName string, projectDir string) {
	CommandStartAndWait(ctx.w, projectDir, "C:\\Windows\\System32\\cmd.exe", "/c", "func", "azure", "functionapp", "publish", funcName)
}
