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
	Pos  *parser.Pos
}

// ASTAssign is an assignment statement.
type ASTAssign struct {
	LHS []ASTNode
	RHS []ASTNode
	Pos *parser.Pos
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
			ast.logExpectIdent()
			return nil
		}
		name := ast.s.Accept()

		if ast.s.AcceptOp("=") {
			p := ast.s.Pos()

			expr := ast.parseExpr()
			if expr == nil {
				return nil
			}

			return &ASTVarDecl{
				Name: name,
				Expr: expr,
				Pos:  p,
			}
		}

		return &ASTVarDecl{
			Name: name,
			Pos:  ast.s.Pos(),
		}
	} else {
		exprs := ast.parseExprList()

		if ast.s.AcceptOp("=") {
			p := ast.s.Pos()
			rhs := ast.parseExprList()
			return &ASTAssign{
				LHS: exprs,
				RHS: rhs,
				Pos: p,
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
		if n.Expr == nil {
			p.Printf("var %s", n.Name.Lit)
		} else {
			p.Printf("var %s = ", n.Name.Lit)
			PrintExpr(p, n.Expr)
		}
	case *ASTExprStmt:
		printExprList(p, n.Exprs)
	default:
		p.Print("/* invalid statement */")
	}
}
