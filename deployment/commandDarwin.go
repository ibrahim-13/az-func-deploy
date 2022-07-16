package deployment

import (
	"io"
)

type typeCommandDarwin struct {
	w io.Writer
}

func (ctx *typeCommandDarwin) DotNetBuild(projectDir string) {
	panic("Not implemented")
}

func (ctx *typeCommandDarwin) ZipBuildOutput(outputDir string, projectDir string) string {
	panic("Not implemented")
}

func (ctx *typeCommandDarwin) AzureZipDeploy(resourceGroup string, funcName string, projectDir string, zipFile string) {
	panic("Not implemented")
}

func (ctx *typeCommandDarwin) AzureFuncZipDeploy(resourceGroup string, funcName string, projectDir string, zipFile string) {
	panic("Not implemented")
}

func (ctx *typeCommandDarwin) FuncDeployProject(funcName string, projectDir string) {
	panic("Not implemented")
}
