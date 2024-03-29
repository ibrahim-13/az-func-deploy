package deployment

import (
	"az-func-deploy/logger"
	"io"
	"path/filepath"
)

type cmdCtxWindows struct {
	w      io.Writer
	cmdExe string
}

func (ctx *cmdCtxWindows) PrintBinaryVersions(logger *logger.Logger) {
	logger.WhiteBlueln("dotnet --version")
	CommandStartAndWait(ctx.w, "", ctx.cmdExe, "/c", "dotnet", "--version")
	logger.WhiteBlueln("az version --output tsv")
	CommandStartAndWait(ctx.w, "", ctx.cmdExe, "/c", "az", "version", "--output", "tsv")
	logger.WhiteBlueln("func --version")
	CommandStartAndWait(ctx.w, "", ctx.cmdExe, "/c", "dotnet", "--version")
}

func (ctx *cmdCtxWindows) DotNetBuild(projectDir string) bool {
	return CommandStartAndWait(ctx.w,
		projectDir,
		ctx.cmdExe,
		"/c",
		"dotnet",
		"build",
		"--configuration",
		"Release")
}

func (ctx *cmdCtxWindows) ZipBuildOutput(zipFile string, projectDir string) bool {
	buildDir := filepath.Join(projectDir, "bin", "Release", "net6.0")
	ok := CommandStartAndWait(ctx.w,
		buildDir,
		ctx.cmdExe,
		"/c",
		"powershell",
		"-command",
		"Compress-Archive -Path * -DestinationPath "+zipFile)
	return ok
}

func (ctx *cmdCtxWindows) AzureZipDeploy(resourceGroup string, funcName string, projectDir string, zipFile string) bool {
	return CommandStartAndWait(ctx.w,
		projectDir,
		ctx.cmdExe,
		"/c",
		"az",
		"webapp",
		"deploy",
		"--resource-group",
		resourceGroup,
		"--name",
		funcName,
		"--src-path",
		zipFile)
}

func (ctx *cmdCtxWindows) AzureFuncZipDeploy(resourceGroup string, funcName string, projectDir string, zipFile string) bool {
	return CommandStartAndWait(ctx.w,
		projectDir,
		ctx.cmdExe,
		"/c",
		"az",
		"functionapp",
		"deployment",
		"source",
		"config-zip",
		"--resource-group",
		resourceGroup,
		"--name",
		funcName,
		"--src",
		zipFile)
}

func (ctx *cmdCtxWindows) FuncDeployProject(funcName string, projectDir string) bool {
	return CommandStartAndWait(ctx.w,
		projectDir,
		ctx.cmdExe,
		"/c",
		"func",
		"azure",
		"functionapp",
		"publish",
		funcName)
}
