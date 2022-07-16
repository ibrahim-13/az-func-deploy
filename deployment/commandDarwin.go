package deployment

import (
	"az-func-deploy/logger"
	"io"
)

type cmdCtxDarwin struct {
	w io.Writer
}

func (ctx *cmdCtxDarwin) PrintBinaryVersions(logger *logger.Logger) {
	panic("Not implemented")
}

func (ctx *cmdCtxDarwin) DotNetBuild(projectDir string) {
	panic("Not implemented")
}

func (ctx *cmdCtxDarwin) ZipBuildOutput(outputDir string, projectDir string) string {
	panic("Not implemented")
}

func (ctx *cmdCtxDarwin) AzureZipDeploy(resourceGroup string, funcName string, projectDir string, zipFile string) {
	panic("Not implemented")
}

func (ctx *cmdCtxDarwin) AzureFuncZipDeploy(resourceGroup string, funcName string, projectDir string, zipFile string) {
	panic("Not implemented")
}

func (ctx *cmdCtxDarwin) FuncDeployProject(funcName string, projectDir string) {
	panic("Not implemented")
}
