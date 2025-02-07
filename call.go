package golang

import "fmt"

type Call struct {
	Fn   *Ident
	Args []Expression
}

func (c *Call) String(imports *ImportMap) string {
	args := ""
	for i, arg := range c.Args {
		if i > 0 {
			args += ", "
		}
		args += arg.String(imports)
	}
	return fmt.Sprintf("%s(%s)", c.Fn.String(imports), args)
}
