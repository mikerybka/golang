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

	Methods []Func
}

func (t *Type) Name() string {
	if t.IsScalar {
		return t.Scalar.Name
	}
	if t.IsPointer {
		return t.Pointer.Name
	}
	if t.IsMap {
		return t.Map.Name
	}
	if t.IsArray {
		return t.Array.Name
	}
	if t.IsStruct {
		return t.Struct.Name
	}
	panic("bad type")
}

func (t *Type) String(imports *ImportMap) string {
	if t.IsScalar {
		return t.Scalar.String(imports)
	} else if t.IsPointer {
		return t.Pointer.String(imports)
	} else if t.IsMap {
		return t.Map.String(imports)
	} else if t.IsArray {
		return t.Array.String(imports)
	} else if t.IsStruct {
		return t.Struct.String(imports)
	} else {
		panic("invalid type")
	}
}
