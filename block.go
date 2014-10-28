package xlang

// block, stmt and tokNode defines the internal memory representation
// of a block, a statement and an entry

// Block defines a series of statements
type Block []Stmt

// Stmt defines a series of token entries
type Stmt []*Entry

// Entry defines a token entry, which could be a single token or a block
type Entry struct {
	Tok   *Tok
	Block Block
}
