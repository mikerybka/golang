package golang

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/mikerybka/util"
)

type Workspace struct {
	Dir string
}

func (w *Workspace) initialized() bool {
	path := filepath.Join(w.Dir, "go.work")
	return util.Exists(path)
}

func (w *Workspace) Init() error {
	if w.initialized() {
		return nil
	}
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
