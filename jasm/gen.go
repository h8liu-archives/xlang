package jasm

import (
	"io"

	"github.com/h8liu/xlang/parser"
)

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

type Object struct {
}
