package lexer

type block []stmt   // a block is a series of statements
type stmt []tokNode // a statement is a series of token nodes
type tokNode struct {
	tok   *Tok
	block block
}

type Block interface {
	Scan() bool
	Stmt() Stmt
}

type Stmt interface {
	Scan() bool
	IsBlock() bool
	Block() Block
	Tok() *Tok
}
