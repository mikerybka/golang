package golang

import (
	"os/exec"

	"github.com/mikerybka/util"
)

func Install(pkg string) error {
	cmd := exec.Command("go", "install", pkg)
	return util.Run(cmd)
}
