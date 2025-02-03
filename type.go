package golang

type Type struct {
	IsScalar bool
	Scalar   *ScalarType

	IsPointer bool
	Pointer   *PointerType

	IsMap bool
	Map   *MapType

	IsArray bool
	Array   *ArrayType

	IsStruct bool
	Struct   *StructType
}

func (t *Type) String(name string, imports *ImportMap) string {
	if t.IsScalar {
		return t.Scalar.String(name, imports)
	} else if t.IsPointer {
		return t.Pointer.String(name, imports)
	} else if t.IsMap {
		return t.Map.String(name, imports)
	} else if t.IsArray {
		return t.Array.String(name, imports)
	} else if t.IsStruct {
		return t.Struct.String(name, imports)
	} else {
		panic("invalid type")
	}
}
