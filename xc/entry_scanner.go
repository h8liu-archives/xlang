package xc

import (
	"github.com/h8liu/xlang/parser"
)

type EntryScanner struct {
}

func NewEntryScanner(s parser.Stmt) *EntryScanner {
	ret := new(EntryScanner)
	return ret
}

func (s *EntryScanner) See(types ...parser.Type) bool {
	panic("todo")
}

func (s *EntryScanner) SeeOp(ops ...string) bool {
	panic("todo")
}

func (s *EntryScanner) AcceptOp(op string) bool {
	panic("todo")
}

func (s *EntryScanner) Accept() *parser.Tok {
	panic("todo")
}

func (s *EntryScanner) Tok() *parser.Tok {
	panic("todo")
}
