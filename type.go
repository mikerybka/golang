package golang

type Type struct {
	Name string

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
