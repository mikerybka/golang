package golang

type Decl struct {
	Name    string
	IsVar   bool
	Var     *Var
	IsConst bool
	Const   *Const
	IsType  bool
	Type    *Type
	IsFunc  bool
	Func    *Func
}
