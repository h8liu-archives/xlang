package parser

import (
	"fmt"
)

type Error struct {
	*Pos
	S string
}

func NewError(p *Pos, s string) *Error {
	return &Error{p, s}
}

func NewBareError(s string) *Error {
	return &Error{nil, s}
}

func (e *Error) Error() string {
	if e.Pos != nil {
		return fmt.Sprintf("%s: %s", e.Pos, e.S)
	}

	return e.S
}
