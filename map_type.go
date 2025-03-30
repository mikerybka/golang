package golang

import "fmt"

type MapType struct {
	Name      string
	KeyType   *Ident
	ValueType *Ident
}

func (t *MapType) String(imports *ImportMap) string {
	return fmt.Sprintf("type %s map[%s]%s\n", t.Name, t.KeyType.String(imports), t.ValueType.String(imports))
}
