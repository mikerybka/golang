package golang

import "fmt"

type PointerType struct {
	UnderlyingType *Ident
}

func (t *PointerType) String(name string, imports *ImportMap) string {
	return fmt.Sprintf("type %s *%s\n", name, t.UnderlyingType.String(imports))
}
