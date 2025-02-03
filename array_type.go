package golang

type ArrayType struct {
	ElementType *Ident
}

func (t *ArrayType) String(name string, imports *ImportMap) string
