package xc

import (
	"bytes"
	"fmt"

	"github.com/h8liu/xlang/parser"
	"github.com/h8liu/xlang/prt"
)

// ASTInfo is just a generic slot for storing syntax info.
type ASTInfo interface{}

// ASTOpExpr describes an expression with a binary operation.
type ASTOpExpr struct {
	A  ASTNode // when A is nil, it is a unary expr
	Op *parser.Tok
	B  ASTNode

	Info ASTInfo
}

// ASTCall describes a function call.
type ASTCall struct {
	Func   ASTNode
	Paras  []ASTNode
	Lparen *parser.Tok
	Rparen *parser.Tok

	Info ASTInfo
}

func (ast *AST) parsePrimaryExpr() ASTNode {
	ret := ast.parseOperand()

	for {
		if ast.s.AcceptOp("(") {
			lp := ast.s.Accepted()

			if ast.s.AcceptOp(")") {
				rp := ast.s.Accepted()
				ret = &ASTCall{
					Func:   ret,
					Paras:  make([]ASTNode, 0),
					Lparen: lp,
					Rparen: rp,
				}
				continue
			}

			lst := ast.parseExprList()
			if lst == nil {
				return nil
			}
			if !ast.expectOp(")") {
				return nil
			}

			rp := ast.s.Accepted()
			ret = &ASTCall{
				Func:   ret,
				Paras:  lst,
				Lparen: lp,
				Rparen: rp,
			}
		} else {
			// done with all those suffixes
			break
		}
	}
	return ret
}

func (ast *AST) parseExprList() []ASTNode {
	ret := make([]ASTNode, 0, 8)
	for {
		expr := ast.parseExpr()
		if expr == nil {
			return nil
		}

		ret = append(ret, expr)

		if !ast.s.AcceptOp(",") {
			break
		}
	}

	return ret
}

func (ast *AST) parseUnaryExpr() ASTNode {
	if ast.s.SeeOp("+", "-") {
		tok := ast.s.Accept()
		x := ast.parseUnaryExpr()
		if x == nil {
			return nil
		}

		return &ASTOpExpr{Op: tok, B: x}
	}

	return ast.parsePrimaryExpr()
}

func (ast *AST) parseBinaryExpr(prec int) ASTNode {
	ret := ast.parseUnaryExpr()
	if ret == nil {
		return nil // error encountered
	}

	if prec == 0 {
		for ast.s.SeeOp("+", "-") {
			tok := ast.s.Accept()
			bop := new(ASTOpExpr)
			bop.A = ret
			bop.Op = tok
			bop.B = ast.parseBinaryExpr(prec + 1)
			ret = bop
		}
	}

	return ret
}

func (ast *AST) parseExpr() ASTNode {
	return ast.parseBinaryExpr(0)
}

func (ast *AST) parseOperand() ASTNode {
	if ast.s.See(parser.TypeIdent, parser.TypeInt) {
		return ast.s.Accept()
	} else if ast.s.AcceptOp("(") {
		ret := ast.parseBinaryExpr(0)
		if !ast.expectOp(")") {
			return nil
		}
		return ret
	}

	ast.logErr(ast.s.Pos(), "expect an operand")
	return nil
}

func (ast *AST) expectOp(op string) bool {
	if ast.s.AcceptOp(op) {
		return true
	}

	ast.errs.Log(ast.s.Pos(), fmt.Sprintf("expect operator '%s'", op))
	return false
}

func (ast *AST) logExpectIdent() {
	ast.errs.Log(ast.s.Pos(), fmt.Sprintf("expect identifier"))
}

func (ast *AST) logErr(pos *parser.Pos, s string) {
	ast.errs.Log(pos, s)
}

// ExprStr returns the string representation of the expression.
// It reflects the tree structure of the expression tree.
func ExprStr(node ASTNode) string {
	buf := new(bytes.Buffer)
	p := prt.New(buf)
	PrintExpr(p, node)
	return buf.String()
}

func printExprList(p *prt.Printer, lst []ASTNode) {
	for i, para := range lst {
		if i != 0 {
			p.Print(",")
		}
		PrintExpr(p, para)
	}
}

func printIdentList(p *prt.Printer, lst []*parser.Tok) {
	for i, t := range lst {
		if i != 0 {
			p.Print(",")
		}
		p.Print(t.Lit)
	}
}

// PrintExpr prints the expression with the printer.
func PrintExpr(p *prt.Printer, node ASTNode) {
	switch n := node.(type) {
	case *ASTOpExpr:
		p.Print("(")
		if n.A != nil {
			PrintExpr(p, n.A)
		}
		p.Print(n.Op.Lit)
		PrintExpr(p, n.B)
		p.Print(")")

	case *ASTCall:
		PrintExpr(p, n.Func)
		p.Print("(")
		printExprList(p, n.Paras)
		p.Print(")")

	case *parser.Tok:
		p.Print(n.Lit)

	default:
		p.Print("(not-a-expr)")
	}
}
