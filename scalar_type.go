package golang

import "fmt"

type ScalarType struct {
	UnderlyingType *Ident
}

func (t *ScalarType) String(name string, imports *ImportMap) string {
	return fmt.Sprintf("type %s %s\n", name, t.UnderlyingType.String(imports))
}
