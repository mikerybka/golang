package golang

type PointerType struct {
	UnderlyingType *Ident
}

func (t *PointerType) String(name string, imports *ImportMap) string
