package golang

type ScalarType struct {
	UnderlyingType *Ident
}

func (t *ScalarType) String(name string, imports *ImportMap) string
