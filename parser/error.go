package parser

import (
	"fmt"
)

// Error defines a compiler parsing error.
type Error struct {
	*Pos
	S string
}

// Error returns the error string.
func (e *Error) Error() string {
	if e.Pos != nil {
		return fmt.Sprintf("%s: %s", e.Pos, e.S)
	}

	return e.S
}
