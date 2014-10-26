package parser

import (
	"fmt"
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

func (p *Pos) StrRowOnly() string {
	return fmt.Sprintf("%s:%d", p.File, p.Row)
}
