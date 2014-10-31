package xc

import (
	"bytes"

	"github.com/h8liu/xlang/parser"
	"github.com/h8liu/xlang/prt"
)

// ASTVarDecl variable declaration statement.
type ASTVarDecl struct {
	Name *parser.Tok
	Expr ASTNode
}

// ASTAssign is an assignment statement.
type ASTAssign struct {
	LHS []ASTNode
	RHS []ASTNode
}

// ASTExprStmt is an expression statement.
type ASTExprStmt struct {
	Exprs []ASTNode
}

func (ast *AST) parseStmt() ASTNode {
	if ast.s.SeeBlock() {
		panic("todo: parsing a block statement")
	} else if ast.s.AcceptKeyword("var") {
		if ast.s.SeeBlock() {
			panic("todo: parsing var decl block")
		}

		// TODO: ident list
		if !ast.s.SeeIdent() {
			return nil
		}
		name := ast.s.Accept()
		if !ast.expectOp("=") {
			return nil
		}
		expr := ast.parseExpr()
		if expr == nil {
			return nil
		}

		return &ASTVarDecl{
			Name: name,
			Expr: expr,
		}
	} else {
		exprs := ast.parseExprList()

		if ast.s.AcceptOp("=") {
			rhs := ast.parseExprList()
			return &ASTAssign{
				LHS: exprs,
				RHS: rhs,
			}
		}

		return &ASTExprStmt{Exprs: exprs}
	}
}

// StmtStr returns the statement representation of a statement.
func StmtStr(node ASTNode) string {
	buf := new(bytes.Buffer)
	p := prt.New(buf)
	PrintStmt(p, node)
	return buf.String()
}

// PrintStmt prints the statement string with the printer.
func PrintStmt(p *prt.Printer, node ASTNode) {
	switch n := node.(type) {
	case *ASTAssign:
		printExprList(p, n.LHS)
		p.Print(" = ")
		printExprList(p, n.RHS)
	case *ASTVarDecl:
		p.Printf("var %s = ", n.Name.Lit)
		PrintExpr(p, n.Expr)
	}
}
