package xast

import (
	"bytes"

	"github.com/h8liu/xlang/parser"
	"github.com/h8liu/xlang/prt"
)

// VarDecl variable declaration statement.
type VarDecl struct {
	Names []*parser.Tok
	Exprs []Node
	Pos   *parser.Pos
}

// Assign is an assignment statement.
type Assign struct {
	LHS []Node
	RHS []Node
	Pos *parser.Pos
}

// ExprStmt is an expression statement.
type ExprStmt struct {
	Expr Node
}

func (t *Tree) parseIdentList() []*parser.Tok {
	var ret []*parser.Tok
	for {
		if !t.s.SeeIdent() {
			t.logExpectIdent()
			return nil
		}
		tok := t.s.Accept()
		ret = append(ret, tok)

		if !t.s.AcceptOp(",") {
			break
		}
	}

	return ret
}

func (t *Tree) parseStmt() Node {
	if t.s.SeeBlock() {
		panic("todo: parsing a block statement")
	} else if t.s.AcceptKeyword("var") {
		if t.s.SeeBlock() {
			panic("todo: parsing var decl block")
		}

		idents := t.parseIdentList()
		if idents == nil {
			return nil
		}

		if t.s.AcceptOp("=") {
			p := t.s.Pos()

			exprs := t.parseExprList()
			if exprs == nil {
				return nil
			}

			return &VarDecl{
				Names: idents,
				Exprs: exprs,
				Pos:   p,
			}
		}

		return &VarDecl{
			Names: idents,
			Pos:   t.s.Pos(),
		}
	} else {
		exprs := t.parseExprList()
		if exprs == nil {
			return nil
		}

		if t.s.AcceptOp("=") {
			p := t.s.Pos()
			rhs := t.parseExprList()
			return &Assign{
				LHS: exprs,
				RHS: rhs,
				Pos: p,
			}
		} else if len(exprs) > 1 {
			t.errs.Log(t.s.Pos(), "unexpected end of statement, expect = or :=")
			return nil
		}

		return &ExprStmt{Expr: exprs[0]}
	}
}

// StmtStr returns the statement representation of a statement.
func StmtStr(node Node) string {
	buf := new(bytes.Buffer)
	p := prt.New(buf)
	PrintStmt(p, node)
	return buf.String()
}

// PrintStmt prints the statement string with the printer.
func PrintStmt(p *prt.Printer, node Node) {
	switch n := node.(type) {
	case *Assign:
		printExprList(p, n.LHS)
		p.Print(" = ")
		printExprList(p, n.RHS)
	case *VarDecl:
		p.Printf("var ")
		printIdentList(p, n.Names)
		if n.Exprs != nil {
			p.Printf(" = ")
			printExprList(p, n.Exprs)
		}
	case *ExprStmt:
		PrintExpr(p, n.Expr)
	default:
		p.Print("/* invalid statement */")
	}
}
