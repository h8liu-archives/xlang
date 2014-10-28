package xc

import (
	"github.com/h8liu/xlang/parser"
)

// AST presents the xlang abstract syntax tree.
type AST struct {
	nodes interface{}
}

func newAST() *AST {
	ret := new(AST)
	return ret
}

func (ast *AST) addStmt(s parser.Stmt) {
	if len(s) == 0 {
		return // empty statement
	}

	lead := s[0]

	if lead.Block != nil {
		panic("todo: parsing a block statement")
	} else {
		t := lead.Tok
		if t == nil {
			panic("empty entry")
		}

		if t.Type == parser.TypeKeyword {
			if t.Lit == "var" {
				// parsing var
			}
		} else {
			// parsing left hand expression
		}
	}
}

func (ast *AST) buildFunc() {
}
