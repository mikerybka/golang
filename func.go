package golang

import (
	"bytes"
	"fmt"
)

type Func struct {
	Name    string
	Recv    *Field
	Inputs  []Field
	Outputs []Field
	Body    []Stmt
}

func (f *Func) String(imports *ImportMap) string {
	buf := bytes.NewBuffer(nil)
	inputs := ""
	for i, in := range f.Inputs {
		if i > 0 {
			inputs += ", "
		}
		inputs += fmt.Sprintf("%s %s", in.Name, in.Type.String(imports))
	}
	fmt.Fprintf(buf, "func %s(%s)", f.Name, inputs)
	if len(f.Outputs) > 0 {
		outputs := "("
		for i, out := range f.Outputs {
			if i > 0 {
				outputs += ", "
			}
			outputs += fmt.Sprintf("%s %s", out.Name, out.Type.String(imports))
		}
		outputs += ")"
		fmt.Fprintf(buf, " %s", f.Outputs[0].Type.Name())
	}
	fmt.Fprintf(buf, "}\n")
	return buf.String()
}
