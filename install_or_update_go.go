package golang

import "github.com/mikerybka/util"

func InstallOrUpdateGo() (bool, error) {
	installedVersion, err := util.GetInstalledGoVersion()
	if err != nil {
		if err.Error() == "exec: \"go\": executable file not found in $PATH" {
			return true, InstallGo()
		}
		return false, err
	}
	latestVersion, err := util.GetLatestGoVersion()
	if err != nil {
		return false, err
	}
	if installedVersion != latestVersion {
		return true, UpdateGo()
	}
	return false, nil
}
