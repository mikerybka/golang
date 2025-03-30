package golang

import "fmt"

type Var struct {
	Name  string
	Type  *Ident
	Value *Expression
}

func (v *Var) String(imports *ImportMap) string {
	return fmt.Sprintf("var %s %s = %s\n", v.Name, v.Type.String(imports), v.Value.String(imports))
}
