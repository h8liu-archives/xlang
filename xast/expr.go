package xast

import (
	"bytes"
	"fmt"

	"github.com/h8liu/xlang/parser"
	"github.com/h8liu/xlang/prt"
)

// OpExpr describes an expression with a binary operation.
type OpExpr struct {
	A  Node // when A is nil, it is a unary expr
	Op *parser.Tok
	B  Node
}

// Call describes a function call.
type Call struct {
	Func   Node
	Paras  []Node
	Lparen *parser.Tok
	Rparen *parser.Tok
}

func (t *Tree) parsePrimaryExpr() Node {
	ret := t.parseOperand()

	for {
		if t.s.AcceptOp("(") {
			lp := t.s.Accepted()

			if t.s.AcceptOp(")") {
				rp := t.s.Accepted()
				ret = &Call{
					Func:   ret,
					Paras:  make([]Node, 0),
					Lparen: lp,
					Rparen: rp,
				}
				continue
			}

			lst := t.parseExprList()
			if lst == nil {
				return nil
			}
			if !t.expectOp(")") {
				return nil
			}

			rp := t.s.Accepted()
			ret = &Call{
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

func (t *Tree) parseExprList() []Node {
	ret := make([]Node, 0, 8)
	for {
		expr := t.parseExpr()
		if expr == nil {
			return nil
		}

		ret = append(ret, expr)

		if !t.s.AcceptOp(",") {
			break
		}
	}

	return ret
}

func (t *Tree) parseUnaryExpr() Node {
	if t.s.SeeOp("+", "-") {
		tok := t.s.Accept()
		x := t.parseUnaryExpr()
		if x == nil {
			return nil
		}

		return &OpExpr{Op: tok, B: x}
	}

	return t.parsePrimaryExpr()
}

func (t *Tree) parseBinaryExpr(prec int) Node {
	ret := t.parseUnaryExpr()
	if ret == nil {
		return nil // error encountered
	}

	if prec == 0 {
		for t.s.SeeOp("+", "-") {
			tok := t.s.Accept()
			bop := new(OpExpr)
			bop.A = ret
			bop.Op = tok
			bop.B = t.parseBinaryExpr(prec + 1)
			ret = bop
		}
	}

	return ret
}

func (t *Tree) parseExpr() Node {
	return t.parseBinaryExpr(0)
}

func (t *Tree) parseOperand() Node {
	if t.s.See(parser.TypeIdent, parser.TypeInt) {
		return t.s.Accept()
	} else if t.s.AcceptOp("(") {
		ret := t.parseBinaryExpr(0)
		if !t.expectOp(")") {
			return nil
		}
		return ret
	}

	t.logErr(t.s.Pos(), "expect an operand")
	return nil
}

func (t *Tree) expectOp(op string) bool {
	if t.s.AcceptOp(op) {
		return true
	}

	t.errs.Log(t.s.Pos(), fmt.Sprintf("expect operator '%s'", op))
	return false
}

func (t *Tree) logExpectIdent() {
	t.errs.Log(t.s.Pos(), fmt.Sprintf("expect identifier"))
}

func (t *Tree) logErr(pos *parser.Pos, s string) {
	t.errs.Log(pos, s)
}

// ExprStr returns the string representation of the expression.
// It reflects the tree structure of the expression tree.
func ExprStr(node Node) string {
	buf := new(bytes.Buffer)
	p := prt.New(buf)
	PrintExpr(p, node)
	return buf.String()
}

func printExprList(p *prt.Printer, lst []Node) {
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
func PrintExpr(p *prt.Printer, node Node) {
	switch n := node.(type) {
	case *OpExpr:
		p.Print("(")
		if n.A != nil {
			PrintExpr(p, n.A)
		}
		p.Print(n.Op.Lit)
		PrintExpr(p, n.B)
		p.Print(")")

	case *Call:
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
