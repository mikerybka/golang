package golang

import "github.com/mikerybka/util"

func InstallOrUpdateGo() (bool, error) {
	installedVersion, err := util.GetInstalledGoVersion()
	if err != nil {
		return true, InstallGo()
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
