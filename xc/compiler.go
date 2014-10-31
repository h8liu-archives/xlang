package xc

import (
	"io"
	"io/ioutil"
	"strings"

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

// NewStrSource creates a source file from a string.
func NewStrSource(f, s string) *Source {
	return &Source{
		File:   f,
		Reader: ioutil.NopCloser(strings.NewReader(s)),
	}
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

	ast := newStmtsAST(block)
	if ast.errs.Len() != 0 {
		return nil, ast.errs
	}

	ast.buildFunc()

	return nil, nil // TODO
}

// BuildExprsAST builds the source as an .xexpr file, where
// each statement is an expression.
// This is useful for testing expression parsing.
func (s *Source) BuildExprsAST() (*ASTBlock, *parser.ErrList) {
	block, errs := parser.Parse(s.File, s.Reader)
	if errs != nil {
		return nil, errs
	}

	ast := newExprsAST(block)
	if ast.errs.Len() != 0 {
		return nil, ast.errs
	}

	return ast.root.(*ASTBlock), nil
}

// BuildStmtsAST builds the source as an .xstmt file, where
// the file works like a function body.
func (s *Source) BuildStmtsAST() (*ASTBlock, *parser.ErrList) {
	block, errs := parser.Parse(s.File, s.Reader)
	if errs != nil {
		return nil, errs
	}

	ast := newStmtsAST(block)
	if ast.errs.Len() != 0 {
		return nil, ast.errs
	}

	return ast.root.(*ASTBlock), nil
}
