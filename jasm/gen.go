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

type builder struct {
	errs *parser.ErrList
}

func newBuilder() *builder {
	ret := new(builder)
	ret.errs = parser.NewErrList()
	return ret
}

func (b *builder) log(p *parser.Pos, f string, args ...interface{}) {
	b.errs.Log(p, f, args...)
}

func (b *builder) buildTops(stmts []*parser.Stmt) {
	for i, stmt := range stmts {
		if len(stmt.Entries) == 0 {
			b.log(stmt.End.Pos, "unexpected top-level empty statement")
			continue
		}

		lead := stmt.Entries[0]
		if lead.Block != nil {
			b.log(lead.Block.Lbrace.Pos, "unexpected block")
		}

		tok := lead.Tok
		if tok.Type != parser.TypeKeyword {
			b.log(tok.Pos, "expect a keyword for leading")
			continue
		}

		switch tok.Lit {
		case "import":
			if i != 0 {
				b.log(tok.Pos, "import must be the first top-level stmt")
			}
		case "func":
			panic("todo")
		case "var":
			panic("todo")
		case "const":
			panic("todo")
		}
	}
}

func build(b *parser.Block) (*Object, *parser.ErrList) {
	bd := newBuilder()
	bd.buildTops(b.Stmts)

	if bd.errs.Len() > 0 {
		return nil, bd.errs
	}
	panic("todo")
}
