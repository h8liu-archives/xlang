package parser

// block, stmt and tokNode defines the internal memory representation
// of a block, a statement and an entry

type Block []Stmt  // a block is a series of statements
type Stmt []*Entry // a statement is a series of token nodes
type Entry struct {
	// en entry is either a token or a block
	Tok   *Tok
	Block Block
}
