package golang

import (
	"os"
	"os/exec"
)

func Install(pkg string) error {
	cmd := exec.Command("go", "install", pkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
