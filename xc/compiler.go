package xc

import (
	"io"

	"github.com/h8liu/xlang/parser"
)

// Header defines all the constants and symbols.
type Header struct {
}

// Object defines the compiling result of a source file.
type Object struct {
	header *Header
}

// Header returns the header of an object file.
func (obj *Object) Header() *Header {
	return obj.header
}

// PubHeader returns the public header of an object file.
func (obj *Object) PubHeader() *Header {
	return nil
}

// Lib defines the library context of a program.
type Lib struct {
}

// Source defines the context required to compile a single source file.
type Source struct {
	File   string        // the file path
	Reader io.ReadCloser // reader for the file content

	Lib *Lib // dependency library
}

// Compile compiles a source file into an object or report the errors.
func (s *Source) Compile() (*Object, *parser.ErrList) {
	return nil, nil
}

// CompileFunc treats the file as the body of the main function.
// It equivalent as wrapping the body inside a `func main() { }`.
func (s *Source) CompileFunc() (*Object, *parser.ErrList) {
	block, errs := parser.Parse(s.File, s.Reader)
	if errs != nil {
		return nil, errs
	}

	tree := newBlockAST(block)
	tree.buildFunc()

	return nil, nil // TODO
}

func buildAST(b parser.Block) {
}