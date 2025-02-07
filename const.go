package golang

import "fmt"

type Const struct {
	Type  *Ident
	Value *Expression
}

func (c *Const) String(name string, imports *ImportMap) string {
	return fmt.Sprintf("const %s %s = %s\n", name, c.Type.String(imports), c.Value.String(imports))
}
