package golang

type File struct {
	PkgName string
	Imports map[string]string
	Decls   []Decl
}
