package golang

import "fmt"

type ScalarType struct {
	Name           string
	UnderlyingType *Ident
}

func (t *ScalarType) String(imports *ImportMap) string {
	return fmt.Sprintf("type %s %s\n", t.Name, t.UnderlyingType.String(imports))
}
