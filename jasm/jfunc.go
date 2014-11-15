package jasm

import (
	"github.com/h8liu/xlang/parser"
)

// a jasm function, a code block to be compiled
type jfunc struct {
	name   *parser.Tok
	blocks []*bb          // function basic blocks
	labels map[string]int // function basic block labels
	loc    uint32
}

// basic block in a function
type bb struct {
	labels []string // leading names

	insts  []*symInst
	offset uint32
}
