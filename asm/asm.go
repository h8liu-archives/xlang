package asm

import (
	"io"

	"github.com/h8liu/xlang/parser"
)

// Object is an assembly object
type Object struct {
}

func Assemble(f string, r io.ReadCloser) (*Object, *parser.ErrList) {
	block, errs := parser.Parse(f, r)
	if errs != nil {
		return nil, errs
	}

	obj, errs := build(block)
	if errs != nil {
		return nil, errs
	}

	return obj, nil
}

func build(b *parser.Block) (*Object, *parser.ErrList) {
	panic("todo")
}
