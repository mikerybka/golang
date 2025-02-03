package golang

import (
	"fmt"
	"path/filepath"
)

type ImportMap struct {
	namesToPaths map[string]string
	pathsToNames map[string]string
}

// AddImport adds an import path to the import map and returns the
// name that will be used.
func (im *ImportMap) AddImport(path string) string {
	// If we've already imported this path, return the local name.
	if name, ok := im.pathsToNames[path]; ok {
		return name
	}

	// Compute the local name
	name := filepath.Base(path)
	// Add underscores to the name if it's already taken.
	for {
		if _, ok := im.namesToPaths[name]; !ok {
			break
		}
		name += "_"
	}

	// Add the import to the mapping
	im.namesToPaths[name] = path
	im.pathsToNames[path] = name

	return name
}

func (im *ImportMap) LocalName(path string) string {
	return im.pathsToNames[path]
}

func (im *ImportMap) String() string {
	importStrings := []string{}
	for name, path := range im.namesToPaths {
		if name == filepath.Base(path) {
			importString := fmt.Sprintf("\"%s\"", path)
			importStrings = append(importStrings, importString)
		} else {
			importString := fmt.Sprintf("%s \"%s\"", name, path)
			importStrings = append(importStrings, importString)
		}
	}

	switch len(importStrings) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("import \"%s\"\n", importStrings[0])
	default:
		s := "import (\n"
		for _, str := range importStrings {
			s += fmt.Sprintf("\t%s\n", str)
		}
		s += ")\n"
		return s
	}
}
