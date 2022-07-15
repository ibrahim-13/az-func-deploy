package deployment

import (
	"io"
	"os/exec"
)

func CommandStartAndWait(w io.Writer, name string, param ...string) {
	cmd := exec.Command(name, param...)
	cmd.Stdout = w
	cmd.Stderr = w
	cmd.Start()
	cmd.Wait()
}
