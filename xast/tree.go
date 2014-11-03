package xast

import (
	"github.com/h8liu/xlang/parser"
)

// Node basically could be anything
type Node interface{}

// Tree presents the xlang abstract syntax tree.
type Tree struct {
	errs *parser.ErrList
	s    *parser.EntryScanner
}

// Block is a scoped block.
type Block struct {
	Nodes []Node
}

// Module is a module that consists of top-level declarations.
type Module struct {
}

func newTree() *Tree {
	ret := new(Tree)
	ret.errs = parser.NewErrList()

	return ret
}

// NewStmts converts a parser block into a tree of statements.
func NewStmts(b *parser.Block) (*Block, *parser.ErrList) {
	ret := newTree()
	root := new(Block)
	ret.parseStmts(root, b)

	if !ret.errs.Empty() {
		return nil, ret.errs
	}
	return root, nil
}

// NewExprs converts a parser block into a tree of expressions.
func NewExprs(b *parser.Block) (*Block, *parser.ErrList) {
	ret := newTree()
	root := new(Block)
	ret.parseExprs(root, b)

	if !ret.errs.Empty() {
		return nil, ret.errs
	}
	return root, nil
}

func (t *Tree) parseProg(ret *Module, b *parser.Block) {
	panic("todo")
}

func (t *Tree) parseStmts(ret *Block, b *parser.Block) {
	for _, s := range b.Stmts {
		if len(s) == 0 {
			continue // empty statement
		}

		t.s = parser.NewEntryScanner(s)
		expr := t.parseStmt()
		if t.s.Entry() != nil {
			t.errs.Log(t.s.Pos(), "expect end of statement")
			continue
		}

		ret.Nodes = append(ret.Nodes, expr)
	}
}

func (t *Tree) parseExprs(ret *Block, b *parser.Block) {
	for _, s := range b.Stmts {
		if len(s) == 0 {
			continue // empty expr
		}

		t.s = parser.NewEntryScanner(s)
		expr := t.parseExpr()
		if t.s.Entry() != nil {
			t.errs.Log(t.s.Pos(), "expect end of statement")
			continue
		}

		ret.Nodes = append(ret.Nodes, expr)
	}
}
