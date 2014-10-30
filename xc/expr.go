package xc

import (
	"fmt"

	"github.com/h8liu/xlang/parser"
)

// ASTOpExpr describes an expression with a binary operation.
type ASTOpExpr struct {
	A  ASTNode // when A is nil, it is a unary expr
	Op *parser.Tok
	B  ASTNode
}

func (ast *AST) parsePrimaryExpr() ASTNode {
	// might add more stuff here
	return ast.parseOperand()
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

func (ast *AST) parseBinaryExpr() ASTNode {
	ret := ast.parseUnaryExpr()
	if ret == nil {
		return nil // error encountered
	}

	for ast.s.SeeOp("+", "-") {
		tok := ast.s.Accept()
		bop := new(ASTOpExpr)
		bop.A = ret
		bop.Op = tok
		bop.B = ast.parseBinaryExpr()
		ret = bop
	}

	return ret
}

func (ast *AST) parseExpr() ASTNode {
	return ast.parseBinaryExpr()
}

func (ast *AST) parseOperand() ASTNode {
	if ast.s.See(parser.TypeIdent, parser.TypeInt) {
		return ast.s.Accept()
	} else if ast.s.AcceptOp("(") {
		ret := ast.parseBinaryExpr()
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
