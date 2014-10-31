package xc

import (
	"bytes"
	"fmt"

	"github.com/h8liu/xlang/parser"
)

// ASTOpExpr describes an expression with a binary operation.
type ASTOpExpr struct {
	A  ASTNode // when A is nil, it is a unary expr
	Op *parser.Tok
	B  ASTNode
}

type ASTCall struct {
	Func  ASTNode
	Paras []ASTNode
}

func (ast *AST) parsePrimaryExpr() ASTNode {
	ret := ast.parseOperand()

	for {
		if ast.s.AcceptOp("(") {
			lst := ast.parseExprList()
			if lst == nil {
				return nil
			}

			ret = &ASTCall{
				Func:  ret,
				Paras: lst,
			}
			if !ast.expectOp(")") {
				return nil
			}
		} else {
			break
		}
	}
	return ret
}

func (ast *AST) parseExprList() []ASTNode {
	if ast.s.SeeOp(")") {
		return make([]ASTNode, 0)
	}

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

func (ast *AST) logErr(pos *parser.Pos, s string) {
	ast.errs.Log(pos, s)
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
		printExpr(buf, n.B)
		fmt.Fprint(buf, ")")
	case *ASTCall:
		printExpr(buf, n.Func)
		fmt.Fprint(buf, "(")
		for i, p := range n.Paras {
			if i != 0 {
				fmt.Fprint(buf, ",")
			}
			printExpr(buf, p)
		}
		fmt.Fprint(buf, ")")
	case *parser.Tok:
		fmt.Fprint(buf, n.Lit)
	default:
		fmt.Fprint(buf, "(not-a-expr)")
	}
}
