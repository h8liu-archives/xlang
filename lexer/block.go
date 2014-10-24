package lexer

type block []stmt   // a block is a series of statements
type stmt []tokNode // a statement is a series of token nodes
type tokNode struct {
	tok   *Tok
	block block
}

// Block defines the interface to scan a block,
// which is a series of statements.
type Block interface {
	Scan() bool
	Stmt() Stmt
}

// Stmt defines the interface to read a statement
// which is a serios of entries,
// where an entry is either a block or a token.
type Stmt interface {
	Scan() bool
	IsBlock() bool
	Block() Block
	Tok() *Tok
}
