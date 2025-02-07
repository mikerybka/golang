package golang

import (
	"bytes"
	"fmt"
)

type StructType struct {
	Fields []Field
}

func (t *StructType) String(name string, imports *ImportMap) string {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintf(buf, "type %s struct {\n", name)
	for _, f := range t.Fields {
		fmt.Fprintf(buf, "\t%s %s\n", f.Name, f.Type.String(imports))
	}
	fmt.Fprintf(buf, "}\n")
	return buf.String()
}
