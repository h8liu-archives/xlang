package parser

// Type is the type of a token
type Type int

// Available types of a token
const (
	TypeInvalid Type = iota
	TypeComment
	TypeOperator
	TypeIdent
	TypeInt
	TypeFloat
	TypeString
)

// Pos defines the position of a token in a file
type Pos struct {
	File string
	Row  int
	Col  int
}

// Tok is a token in a file
type Tok struct {
	Type Type
	Lit  string
	Pos  *Pos
}

