package golang

import "github.com/mikerybka/util"

func InstallOrUpdateGo() error {
	installedVersion, err := util.GetInstalledGoVersion()
	if err != nil {
		if err.Error() == "exec: \"go\": executable file not found in $PATH" {
			return InstallGo()
		}
		return err
	}
	latestVersion, err := util.GetLatestGoVersion()
	if err != nil {
		return err
	}
	if installedVersion != latestVersion {
		return UpdateGo()
	}
	return nil
}
