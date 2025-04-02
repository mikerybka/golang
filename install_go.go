package golang

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func InstallGo() error {
	// Find dl link
	res, err := http.Get("https://go.dev/dl/?mode=json")
	if err != nil {
		return err
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	releases := []Release{}
	err = json.Unmarshal(b, &releases)
	if err != nil {
		return err
	}

	// Download archive
	releaseFile, err := releases[0].FindFile(runtime.GOOS, runtime.GOARCH)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("https://go.dev/dl/%s", releaseFile.Filename)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected http status %d from %s", res.StatusCode, url)
	}
	archivePath := "/tmp/go.tar.gz"
	f, err := os.Create(archivePath)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, res.Body)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}

	// Extract to /usr/local/go
	cmd := exec.Command("tar", "-C", "/usr/local", "-xzf", archivePath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", cmd, out)
	}

	// Delete archive
	return os.RemoveAll(archivePath)
}
