package golang

import "fmt"

type PointerType struct {
	Name           string
	UnderlyingType *Ident
}

func (t *PointerType) String(imports *ImportMap) string {
	return fmt.Sprintf("type %s *%s\n", t.Name, t.UnderlyingType.String(imports))
}
