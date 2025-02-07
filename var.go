package golang

import "fmt"

type Var struct {
	Type  *Ident
	Value *Expression
}

func (v *Var) String(name string, imports *ImportMap) string {
	return fmt.Sprintf("var %s %s = %s\n", name, v.Type.String(imports), v.Value.String(imports))
}
