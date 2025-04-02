package golang

import "fmt"

type Release struct {
	Version string        `json:"version"`
	Stable  bool          `json:"stable"`
	Files   []ReleaseFile `json:"files"`
}

func (r *Release) FindFile(os, arch string) (*ReleaseFile, error) {
	for _, f := range r.Files {
		if f.OS == os && f.Arch == arch {
			return &f, nil
		}
	}
	return nil, fmt.Errorf("%s-%s file not found", os, arch)
}
