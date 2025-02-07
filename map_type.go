package golang

import "fmt"

type MapType struct {
	KeyType   *Ident
	ValueType *Ident
}

func (t *MapType) String(name string, imports *ImportMap) string {
	return fmt.Sprintf("type %s map[%s]%s\n", name, t.KeyType.String(imports), t.ValueType.String(imports))
}
