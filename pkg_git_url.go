package golang

import (
	"fmt"
	"os"
	"strings"
)

func PkgGitURL(pkg string) (string, error) {
	pkgParts := strings.Split(pkg, "/")
	if len(pkgParts) < 3 {
		return "", os.ErrNotExist
	}
	url := fmt.Sprintf("https://%s/%s/%s.git", pkgParts[0], pkgParts[1], pkgParts[2])
	return url, nil
}
