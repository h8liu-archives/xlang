package xc

import (
	"github.com/h8liu/xlang/parser"
)

type symbol struct {
	name string
	pos  *parser.Pos
	typ  *xtype
	v    *enode
}

type nameIndex struct {
	all map[string]*symbol
}

func (ind *nameIndex) put(s *symbol) bool {
	if s == nil {
		panic("cannot put nil")
	}

	name := s.name
	_, found := ind.all[name]
	if found {
		return false
	}
	ind.all[name] = s
	return true
}

func (ind *nameIndex) find(name string) *symbol {
	return ind.all[name]
}

type scope struct {
	stack []*nameIndex
}

func newScope() *scope {
	ret := new(scope)
	return ret
}

func (s *scope) push() {
	index := new(nameIndex)
	s.stack = append(s.stack, index)
}

func (s *scope) pop() {
	if len(s.stack) == 0 {
		panic("nothing to pop")
	}

	s.stack = s.stack[:len(s.stack)-1]
}

func (s *scope) top() *nameIndex {
	if len(s.stack) == 0 {
		panic("no top yet")
	}

	return s.stack[len(s.stack)-1]
}

func (s *scope) put(sym *symbol) bool {
	top := s.top()
	return top.put(sym)
}

func (s *scope) findTop(name string) *symbol {
	return s.top().find(name)
}

func (s *scope) find(name string) *symbol {
	n := len(s.stack)
	for i := n - 1; i >= 0; i-- {
		ind := s.stack[i]
		ret := ind.find(name)
		if ret != nil {
			return ret
		}
	}

	return nil
}
