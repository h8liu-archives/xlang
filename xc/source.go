package xc

import (
	"io"
	"io/ioutil"
	"strings"

	"github.com/h8liu/xlang/asm"
	"github.com/h8liu/xlang/parser"
	"github.com/h8liu/xlang/xast"
)

// Source defines the context required to compile a single source file.
type Source struct {
	File   string        // the file path
	Reader io.ReadCloser // reader for the file content
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
	tree, errs := s.BuildStmtsAST()
	if errs != nil {
		return nil, errs
	}

	b := newBuilder(tree)
	b.buildFunc()
	if !b.errs.Empty() {
		return nil, b.errs
	}

	return b.obj, nil
}

// BuildExprsAST builds the source as an .xexpr file, where
// each statement is an expression.
// This is useful for testing expression parsing.
func (s *Source) BuildExprsAST() (*xast.Block, *parser.ErrList) {
	block, errs := parser.Parse(s.File, s.Reader)
	if errs != nil {
		return nil, errs
	}

	tree, errs := xast.NewExprs(block)
	if errs != nil {
		return nil, errs
	}

	return tree, nil
}

// BuildStmtsAST builds the source as an .xstmt file, where
// the file works like a function body.
func (s *Source) BuildStmtsAST() (*xast.Block, *parser.ErrList) {
	block, errs := parser.Parse(s.File, s.Reader)
	if errs != nil {
		return nil, errs
	}

	tree, errs := xast.NewStmts(block)
	if errs != nil {
		return nil, errs
	}

	return tree, nil
}

// Assemble assembles a assembly file.
func (s *Source) Assemble() (*asm.Object, *parser.ErrList) {
	ret, errs := asm.Assemble(s.File, s.Reader)
	if errs != nil {
		return nil, errs
	}

	return ret, nil
}
