package golang

type Expression struct {
	IsLiteral bool
	Literal   string

	IsIdent bool
	Ident   *Ident

	IsCall bool
	Call   *Call
}

func (e *Expression) String(imports *ImportMap) string {
	if e.IsLiteral {
		return e.Literal
	}
	if e.IsIdent {
		return e.Ident.String(imports)
	}
	if e.IsCall {
		return e.Call.String(imports)
	}
	panic("bad expression")
}
