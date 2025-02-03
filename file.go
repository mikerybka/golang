package golang

import (
	"fmt"
	"go/format"
)

type File struct {
	PkgName string
	Decls   []Decl
}

func (f *File) String() string {
	imports := &ImportMap{}
	decls := ""
	for _, d := range f.Decls {
		decls += d.String(imports) + "\n"
	}
	s := fmt.Sprintf("package %s\n\n%s%s", f.PkgName, imports.String(), decls)
	b, err := format.Source([]byte(s))
	if err != nil {
		panic(err)
	}
	return string(b)
}
