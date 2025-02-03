package golang

type StructType struct {
	Fields []Field
}

func (t *StructType) String(name string, imports *ImportMap) string
