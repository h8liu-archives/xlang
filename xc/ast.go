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

	root ASTNode
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
		ast.addStmt(ret, s)
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
			ast.errs.Log(ast.s.Pos(), "expression does not end cleanly")
			continue
		}

		ret.Nodes = append(ret.Nodes, expr)
	}
}

func (ast *AST) addStmt(b *ASTBlock, s parser.Stmt) {
	if len(s) == 0 {
		return // empty statement
	}

	ast.s = parser.NewEntryScanner(s)

	if ast.s.SeeBlock() {
		panic("todo: parsing a block statement")
	} else if ast.s.See(parser.TypeKeyword) {
		t := ast.s.Tok()
		if t.Lit == "var" {
			// parsing var
		}
	} else {
		// parsing left hand expression
		_ = ast.parseExpr()
	}
}

func (ast *AST) buildFunc() {

}
