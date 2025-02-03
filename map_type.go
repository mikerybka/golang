package golang

type MapType struct {
	KeyType   *Ident
	ValueType *Ident
}

func (t *MapType) String(name string, imports *ImportMap) string
