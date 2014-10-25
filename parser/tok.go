package parser

import (
	"fmt"
)

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
	TypeKeyword
)


// Pos defines the position of a token in a file
type Pos struct {
	File string
	Row  int
	Col  int
}

func (p *Pos) String() string {
	return fmt.Sprintf("%s:%d:%d", p.File, p.Row, p.Col)
}

// Tok is a token in a file
type Tok struct {
	Type Type
	Lit  string
	Pos  *Pos
}

var typeStr = map[Type]string {
	TypeInvalid: "invalid",
	TypeComment: "comment",
	TypeOperator: "operator",
	TypeIdent: "ident",
	TypeKeyword: "keyword",
	TypeInt: "int",
	TypeFloat: "float",
	TypeString: "string",
}

func (t Type) String() string {
	ret, found := typeStr[t]
	if !found {
		ret = fmt.Sprintf("type-%d", t)
	}
	return ret
}

func (t *Tok) String() string {
	return fmt.Sprintf("%s: <%s> %q",
		t.Pos.String(),
		t.Type.String(),
		t.Lit,
	)
}
