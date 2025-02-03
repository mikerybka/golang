package golang

type Const struct {
	Type  string
	Value string
}

func (c *Const) String(name string, imports *ImportMap) string
