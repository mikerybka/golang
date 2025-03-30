package golang

type Decl struct {
	Name string

	IsVar   bool
	Var     *Var
	IsConst bool
	Const   *Const
	IsType  bool
	Type    *Type
	IsFunc  bool
	Func    *Func
}

func (d *Decl) String(imports *ImportMap) string {
	if d.IsVar {
		return d.Var.String(imports)
	}
	if d.IsConst {
		return d.Const.String(imports)
	}
	if d.IsType {
		return d.Type.String(imports)
	}
	if d.IsFunc {
		return d.Func.String(imports)
	}
	panic("invalid decl")
}
