package lexer

type Type int

type Pos struct {
	File string
	Row  int
	Col  int
}

const (
	TypeIdent Type = iota
	TypeInt
	// TypeFloat
	// TypeString
	TypeComment
	TypeOperator
)

type Tok struct {
	Type Type
	Lit  string
	Pos  *Pos
}

// A program is a block.
// A block is a series of statements.
// A statement is a series of tokens or blocks
//
// A lexer uses a block builder to build a block.
// It calls Token() for every non-block token in the statement
// and calls AddBlock() for adding a block token, which will return a nested
// BlockBuilder.
// It calls EndStmt() at the end of every statement.
// It calls Close() at the end of the block.
type BlockBuilder interface {
	EndStmt()
	AddBlock() BlockBuilder
	AddTok(t *Tok)
	Close()
}
