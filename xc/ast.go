package xc

import (
	"github.com/h8liu/xlang/parser"
)

// ASTNode basically could be anything
type ASTNode interface{}

// ASTBlock presents the xlang abstract syntax tree.
type AST struct {
	errs *parser.ErrList
	root ASTNode
}

// ASTBlock is a scoped block
type ASTBlock struct {
	nodes []ASTNode
}

func newBlockAST(b parser.Block) *AST {
	ret := new(AST)
	ret.errs = parser.NewErrList()
	root := new(ASTBlock)
	ret.parseBlock(root, b)
	ret.root = root

	return ret
}

func (ast *AST) parseBlock(ret *ASTBlock, b parser.Block) {
	for _, s := range b {
		ast.addStmt(ret, s)
	}
}

func (ast *AST) addStmt(b *ASTBlock, s parser.Stmt) {
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
