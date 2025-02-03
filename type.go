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

func (t *Type) String(name string, imports *ImportMap) string
