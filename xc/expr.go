package xc

import (
	"github.com/h8liu/xlang/parser"
)

type ASTOpExpr struct {
	A  ASTNode // when A is nil, it is a unary expr
	Op string
	B  ASTNode
}

func (ast *AST) parsePrimaryExpr(s *EntryScanner) ASTNode {
	// might add more stuff here
	return ast.parseOperand(s)
}

func (ast *AST) parseUnaryExpr(s *EntryScanner) ASTNode {
	if s.SeeOp("+", "-") {
		tok := s.Accept()
		x := ast.parseUnaryExpr(s)
		if x == nil {
			return nil
		}

		return &ASTOpExpr{Op: tok.Lit, B: x}
	}

	return ast.parsePrimaryExpr(s)
}

func (ast *AST) parseBinaryExpr(s *EntryScanner) ASTNode {
	ret := ast.parseUnaryExpr(s)
	if ret == nil {
		return nil // error encountered
	}

	for s.SeeOp("+", "-") {
		tok := s.Accept()
		bop := new(ASTOpExpr)
		bop.A = ret
		bop.Op = tok.Lit
		bop.B = ast.parseBinaryExpr(s)
		ret = bop
	}

	return ret
}

func (ast *AST) parseOperand(s *EntryScanner) ASTNode {
	if s.See(parser.TypeIdent, parser.TypeInt) {
		return s.Accept()
	} else if s.AcceptOp("(") {
		ret := ast.parseBinaryExpr(s)
		if !ast.expectOp(s, ")") {
			return nil
		}
		return ret
	}

	ast.logErr(s.Tok().Pos, "expect an operand")
	return nil
}

func (ast *AST) expectOp(s *EntryScanner, op string) bool {
	panic("todo")
}

func (ast *AST) logErr(pos *parser.Pos, s string) {
	ast.errs.Log(pos, s)
}
