package xc

import (
	"github.com/h8liu/xlang/parser"
)

type symbol struct {
	name string
	pos  *parser.Pos
	v    *enode
}
