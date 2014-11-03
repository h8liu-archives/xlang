package xast

import (
	"github.com/h8liu/xlang/parser"
)

// ASTNode basically could be anything
type Node interface{}

// AST presents the xlang abstract syntax tree.
type Tree struct {
	errs *parser.ErrList
	s    *parser.EntryScanner

	root Node
}

// ASTBlock is a scoped block.
type Block struct {
	Nodes []Node
}

// ASTModule is a module that consists of top-level declarations.
type Module struct {
}

func newTree() *Tree {
	ret := new(Tree)
	ret.errs = parser.NewErrList()

	return ret
}

func NewStmts(b *parser.Block) (*Tree, *parser.ErrList) {
	ret := newTree()
	root := new(Block)
	ret.parseStmts(root, b)
	ret.root = root

	return ret, ret.errs
}

func NewExprs(b *parser.Block) (*Tree, *parser.ErrList) {
	ret := newTree()
	root := new(Block)
	ret.parseExprs(root, b)
	ret.root = root

	return ret, ret.errs
}

func (t *Tree) Root() Node            { return t.root }
func (t *Tree) Errs() *parser.ErrList { return t.errs }

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
