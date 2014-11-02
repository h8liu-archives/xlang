package xc

import (
	"github.com/h8liu/xlang/parser"
)

// ASTNode basically could be anything
type ASTNode interface{}

// AST presents the xlang abstract syntax tree.
type AST struct {
	errs *parser.ErrList
	s    *parser.EntryScanner

	root  ASTNode
	scope *scope
	ir    *irBuilder
	obj   *Object
}

// ASTBlock is a scoped block.
type ASTBlock struct {
	Nodes []ASTNode
}

// ASTModule is a module that consists of top-level declarations.
type ASTModule struct {
}

func newAST() *AST {
	ret := new(AST)
	ret.errs = parser.NewErrList()
	return ret
}

func newStmtsAST(b *parser.Block) *AST {
	ret := newAST()
	root := new(ASTBlock)
	ret.parseStmts(root, b)
	ret.root = root

	return ret
}

func newExprsAST(b *parser.Block) *AST {
	ret := newAST()
	root := new(ASTBlock)
	ret.parseExprs(root, b)
	ret.root = root
	return ret
}

func (ast *AST) parseProg(ret *ASTModule, b *parser.Block) {
	panic("todo")
}

func (ast *AST) parseStmts(ret *ASTBlock, b *parser.Block) {
	for _, s := range b.Stmts {
		if len(s) == 0 {
			continue // empty statement
		}

		ast.s = parser.NewEntryScanner(s)
		expr := ast.parseStmt()
		if ast.s.Entry() != nil {
			ast.errs.Log(ast.s.Pos(), "expect end of statement")
			continue
		}

		ret.Nodes = append(ret.Nodes, expr)
	}
}

func (ast *AST) parseExprs(ret *ASTBlock, b *parser.Block) {
	for _, s := range b.Stmts {
		if len(s) == 0 {
			continue // empty expr
		}

		ast.s = parser.NewEntryScanner(s)
		expr := ast.parseExpr()
		if ast.s.Entry() != nil {
			ast.errs.Log(ast.s.Pos(), "expect end of statement")
			continue
		}

		ret.Nodes = append(ret.Nodes, expr)
	}
}
