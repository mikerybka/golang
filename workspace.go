package golang

import (
	"os"
	"os/exec"
)

type Workspace string

func (w Workspace) Build(pkg string, outFile string) error {
	cmd := exec.Command("go", "build", "-o", outFile, pkg)
	cmd.Dir = string(w)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
