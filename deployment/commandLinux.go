package deployment

import (
	"az-func-deploy/logger"
	"io"
)

type cmdCtxLinux struct {
	w io.Writer
}

func (ctx *cmdCtxLinux) PrintBinaryVersions(logger *logger.Logger) {
	panic("Not implemented")
}

func (ctx *cmdCtxLinux) DotNetBuild(projectDir string) {
	panic("Not implemented")
}

func (ctx *cmdCtxLinux) ZipBuildOutput(outputDir string, projectDir string) string {
	panic("Not implemented")
}

func (ctx *cmdCtxLinux) AzureZipDeploy(resourceGroup string, funcName string, projectDir string, zipFile string) {
	panic("Not implemented")
}

func (ctx *cmdCtxLinux) AzureFuncZipDeploy(resourceGroup string, funcName string, projectDir string, zipFile string) {
	panic("Not implemented")
}

func (ctx *cmdCtxLinux) FuncDeployProject(funcName string, projectDir string) {
	panic("Not implemented")
}
