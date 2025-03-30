package golang

type Package struct {
	Name  string
	Decls []Decl
}

func ReadPackage(path string) (*Package, error)
