package golang

import (
	"fmt"
	"os/exec"
)

type Workspace struct {
	Dir string
}

func (w *Workspace) Init() error {
	cmd := exec.Command("go", "work", "init")
	cmd.Dir = w.Dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s: %s", cmd.String(), err, out)
	}
	return nil
}

func (w *Workspace) Build(pkg string, outFile string) error {
	cmd := exec.Command("go", "build", "-o", outFile, pkg)
	cmd.Dir = w.Dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s: %s", cmd.String(), err, out)
	}
	return nil
}
