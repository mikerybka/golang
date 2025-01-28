package golang

type Stmt struct {
	IsBadStmt bool
	BadStmt   *BadStmt

	IsDeclStmt bool
	DeclStmt   *DeclStmt

	IsEmptyStmt bool
	EmptyStmt   *EmptyStmt

	IsLabeledStmt bool
	LabeledStmt   *LabeledStmt

	IsExprStmt bool
	ExprStmt   *ExprStmt

	IsSendStmt bool
	SendStmt   *SendStmt

	IsIncDecStmt bool
	IncDecStmt   *IncDecStmt

	IsAssignStmt bool
	AssignStmt   *AssignStmt

	IsGoStmt bool
	GoStmt   *GoStmt

	IsDeferStmt bool
	DeferStmt   *DeferStmt

	IsReturnStmt bool
	ReturnStmt   *ReturnStmt

	IsBranchStmt bool
	BranchStmt   *BranchStmt

	IsBlockStmt bool
	BlockStmt   *BlockStmt

	IsIfStmt bool
	IfStmt   *IfStmt

	IsCaseClause bool
	CaseClause   *CaseClause

	IsSwitchStmt bool
	SwitchStmt   *SwitchStmt

	IsTypeSwitchStmt bool
	TypeSwitchStmt   *TypeSwitchStmt

	IsCommClause bool
	CommClause   *CommClause

	IsSelectStmt bool
	SelectStmt   *SelectStmt

	IsForStmt bool
	ForStmt   *ForStmt

	IsRangeStmt bool
	RangeStmt   *RangeStmt
}

type BadStmt struct {
}
type DeclStmt struct {
}
type EmptyStmt struct {
}
type LabeledStmt struct {
}
type ExprStmt struct {
}
type SendStmt struct {
}
type IncDecStmt struct {
}
type AssignStmt struct {
}
type GoStmt struct {
}
type DeferStmt struct {
}
type ReturnStmt struct {
}
type BranchStmt struct {
}
type BlockStmt struct {
}
type IfStmt struct {
}
type CaseClause struct {
}
type SwitchStmt struct {
}
type TypeSwitchStmt struct {
}
type CommClause struct {
}
type SelectStmt struct {
}
type ForStmt struct {
}
type RangeStmt struct {
}
