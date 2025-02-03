package golang

type Var struct {
	Type  Ident
	Value Expression
}

func (v *Var) String(name string, imports *ImportMap) string
