package golang

type Expression struct {
	IsBool bool
	Bool   bool

	IsInt bool
	Int   int

	IsString bool
	String   string

	IsIdent bool
	Ident   *Ident

	IsCall bool
	Call   *Call
}
