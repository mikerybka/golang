package golang

type Func struct {
	Inputs  []Field
	Outputs []Field
	Body    []Stmt
}

func (f *Func) String(name string, imports *ImportMap) string
