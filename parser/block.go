package parser

// block, stmt and tokNode defines the internal memory representation
// of a block, a statement and an entry

type block []stmt // a block is a series of statements
type stmt []entry // a statement is a series of token nodes
type entry struct {
	// en entry is either a token or a block
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
