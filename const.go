package golang

import "fmt"

type Const struct {
	Name  string
	Type  *Ident
	Value *Expression
}

func (c *Const) String(imports *ImportMap) string {
	return fmt.Sprintf("const %s %s = %s\n", c.Name, c.Type.String(imports), c.Value.String(imports))
}
