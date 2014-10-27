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

// Tok is a token in a file
type Tok struct {
	Type Type
	Lit  string
	*Pos
}

var typeStr = map[Type]string{
	TypeInvalid:  "invalid",
	TypeComment:  "comment",
	TypeOperator: "operator",
	TypeIdent:    "ident",
	TypeKeyword:  "keyword",
	TypeInt:      "int",
	TypeFloat:    "float",
	TypeString:   "string",
}

var typeShortStr = map[Type]string{
	TypeInvalid:  "iv",
	TypeComment:  "cm",
	TypeOperator: "op",
	TypeIdent:    "id",
	TypeKeyword:  "kw",
	TypeInt:      "i",
	TypeFloat:    "f",
	TypeString:   "str",
}

func (t Type) String() string {
	ret, found := typeStr[t]
	if !found {
		ret = fmt.Sprintf("type-%d", t)
	}
	return ret
}

func (t Type) ShortStr() string {
	ret, found := typeShortStr[t]
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
