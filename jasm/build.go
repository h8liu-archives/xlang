package jasm

import (
	"github.com/h8liu/xlang/parser"
)

type builder struct {
	s    *parser.EntryScanner
	errs *parser.ErrList

	funcs map[string]*jfunc
}

func newBuilder() *builder {
	ret := new(builder)
	ret.errs = parser.NewErrList()
	ret.funcs = make(map[string]*jfunc)

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

		b.s = parser.NewEntryScanner(stmt)

		if !b.s.See(parser.TypeKeyword) {
			b.log(b.s.Pos(), "expect a keyword for leading")
			continue
		}

		tok := b.s.Tok()
		switch tok.Lit {
		case "import":
			if i != 0 {
				b.log(tok.Pos, "import must be the first top-level stmt")
				continue
			}

			panic("todo")
		case "func":
			f := b.buildFunc()
			if f != nil {
				b.funcs[f.name.Lit] = f
			}
		case "var":
			panic("todo")
		case "const":
			panic("todo")
		}
	}
}

func (b *builder) buildFuncBody(ret *jfunc, body *parser.Block) *jfunc {

	return ret
}

func (b *builder) buildFunc() *jfunc {
	if !b.s.AcceptKeyword("func") {
		panic("must lead with func keyword")
	}

	if !b.s.SeeIdent() {
		b.log(b.s.Pos(), "expect an identifier for function name")
		return nil
	}

	ret := new(jfunc)
	ret.name = b.s.Accept()

	if !b.s.SeeBlock() {
		b.log(b.s.Pos(), "expect the function body")
		return nil
	}

	body := b.s.Block()
	ret = b.buildFuncBody(ret, body)

	if !b.s.End() {
		b.log(b.s.Pos(), "expect end of decl stmt")
		return nil
	}

	return ret
}

func build(b *parser.Block) (*Object, *parser.ErrList) {
	bd := newBuilder()
	bd.buildTops(b.Stmts)

	if bd.errs.Len() > 0 {
		return nil, bd.errs
	}
	panic("todo")
}
