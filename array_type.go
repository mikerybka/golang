package golang

import "fmt"

type ArrayType struct {
	Name        string
	ElementType *Ident
}

func (t *ArrayType) String(imports *ImportMap) string {
	return fmt.Sprintf("type %s []%s\n", t.Name, t.ElementType.String(imports))
}
