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
		return fmt.Sprintf("%s: %s", 
			e.Pos.StrRowOnly(), e.S,
		)
	}

	return fmt.Sprintf("error: %s", e.S)
}

// String returns the error string. It is same as Error().
func (e *Error) String() string {
	return e.Error()
}
