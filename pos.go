package xlang

import (
	"fmt"
)

// Pos defines the position of a token in a file
type Pos struct {
	File string
	Row  int
	Col  int
}

// String returns the string representation of a position
// in the form of "<filename>:<row number>:<col number>"
func (p *Pos) String() string {
	return fmt.Sprintf("%s:%d:%d", p.File, p.Row, p.Col)
}

// StrRowOnly returns the string representation of a poisition
// in the form of "<filename>:<row number>"
func (p *Pos) StrRowOnly() string {
	return fmt.Sprintf("%s:%d", p.File, p.Row)
}
