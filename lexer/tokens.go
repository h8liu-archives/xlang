package lexer

type Type int

type Pos struct {
	File string
	Row  int
	Col  int
}

const (
	TypeIdent Type = iota
	TypeInt
	TypeFloat
	TypeString
	TypeComment
	TypeOperator
)

type Token struct {
	Type Type
	Lit  string
	Pos  *Pos
}

type Block interface {
	EndStmt()
	NewBlock() Block
	WriteToken(t *Token)
	EndBlock()
}
