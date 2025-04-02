package golang

import "os"

func UpdateGo() error {
	err := os.RemoveAll("/usr/local/go")
	if err != nil {
		return err
	}

	return InstallGo()
}
