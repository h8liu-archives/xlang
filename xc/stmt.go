package xc

import (
	"bytes"

	"github.com/h8liu/xlang/parser"
	"github.com/h8liu/xlang/prt"
)

// ASTVarDecl variable declaration statement.
type ASTVarDecl struct {
	Names []*parser.Tok
	Exprs []ASTNode
	Pos   *parser.Pos
}

// ASTAssign is an assignment statement.
type ASTAssign struct {
	LHS []ASTNode
	RHS []ASTNode
	Pos *parser.Pos
}

// ASTExprStmt is an expression statement.
type ASTExprStmt struct {
	Expr ASTNode
}

func (ast *AST) parseIdentList() []*parser.Tok {
	var ret []*parser.Tok
	for {
		if !ast.s.SeeIdent() {
			ast.logExpectIdent()
			return nil
		}
		t := ast.s.Accept()
		ret = append(ret, t)

		if !ast.s.AcceptOp(",") {
			break
		}
	}

	return ret
}

func (ast *AST) parseStmt() ASTNode {
	if ast.s.SeeBlock() {
		panic("todo: parsing a block statement")
	} else if ast.s.AcceptKeyword("var") {
		if ast.s.SeeBlock() {
			panic("todo: parsing var decl block")
		}

		idents := ast.parseIdentList()
		if idents == nil {
			return nil
		}

		if ast.s.AcceptOp("=") {
			p := ast.s.Pos()

			exprs := ast.parseExprList()
			if exprs == nil {
				return nil
			}

			return &ASTVarDecl{
				Names: idents,
				Exprs: exprs,
				Pos:   p,
			}
		}

		return &ASTVarDecl{
			Names: idents,
			Pos:   ast.s.Pos(),
		}
	} else {
		exprs := ast.parseExprList()
		if exprs == nil {
			return nil
		}

		if ast.s.AcceptOp("=") {
			p := ast.s.Pos()
			rhs := ast.parseExprList()
			return &ASTAssign{
				LHS: exprs,
				RHS: rhs,
				Pos: p,
			}
		} else if len(exprs) > 1 {
			ast.errs.Log(ast.s.Pos(), "unexpected end of statement, expect = or :=")
			return nil
		}

		return &ASTExprStmt{Expr: exprs[0]}
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
		p.Printf("var ")
		printIdentList(p, n.Names)
		if n.Exprs != nil {
			p.Printf(" = ")
			printExprList(p, n.Exprs)
		}
	case *ASTExprStmt:
		PrintExpr(p, n.Expr)
	default:
		p.Print("/* invalid statement */")
	}
}
