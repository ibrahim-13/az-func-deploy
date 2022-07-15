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

func CommandAzureZipDeploy(w io.Writer, resourceGroup string, funcName string, projectDir string, zipFile string) {
	CommandStartAndWait(w, projectDir, "C:\\Windows\\System32\\cmd.exe", "/c", "az", "webapp", "deploy", "--resource-group", resourceGroup, "--name", funcName, "--src-path", zipFile)
}
