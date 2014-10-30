package xc

import (
	"bytes"
	"fmt"

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

// ASTBlock is a scoped block
type ASTBlock struct {
	nodes []ASTNode
}

func newBlockAST(b *parser.Block) *AST {
	ret := new(AST)
	ret.errs = parser.NewErrList()
	root := new(ASTBlock)
	ret.parseBlock(root, b)
	ret.root = root

	return ret
}

func (ast *AST) parseBlock(ret *ASTBlock, b *parser.Block) {
	for _, s := range b.Stmts {
		ast.addStmt(ret, s)
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

// ExprStr returns the string representation of the expression.
// It reflects the tree structure of the expression tree.
func ExprStr(node ASTNode) string {
	buf := new(bytes.Buffer)
	printExpr(buf, node)
	return buf.String()
}

func printExpr(buf *bytes.Buffer, node ASTNode) {
	switch n := node.(type) {
	case *ASTOpExpr:
		fmt.Fprint(buf, "(")
		if n.A != nil {
			printExpr(buf, n.A)
		}
		fmt.Fprint(buf, n.Op.Lit)
		fmt.Fprint(buf, n.B)
		fmt.Fprint(buf, ")")
	case *parser.Tok:
		fmt.Fprint(buf, n.Lit)
	}
}
