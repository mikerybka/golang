package golang

import "fmt"

type ArrayType struct {
	ElementType *Ident
}

func (t *ArrayType) String(name string, imports *ImportMap) string {
	return fmt.Sprintf("type %s []%s\n", name, t.ElementType.String(imports))
}
