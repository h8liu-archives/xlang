package parser

// Type is the type of a token
type Type int

// Available types of a token
const (
	TypeInvalid Type = iota
	TypeComment
	TypeOperator
	TypeIdent
	TypeInt
	TypeFloat
	TypeString
)

// Pos defines the position of a token in a file
type Pos struct {
	File string
	Row  int
	Col  int
}

// Tok is a token in a file
type Tok struct {
	Type Type
	Lit  string
	Pos  *Pos
}

// blockBuilder defines the interface to build a block
//
// Conceptually, a block is a series of statements,
// and a statement is a series of tokens or blocks
//
// A lexer uses a block builder to build a block.
// It calls Token() for every non-block token in the statement
// and calls AddBlock() for adding a block token, which will return a nested
// BlockBuilder.
// It calls EndStmt() at the end of every statement.
// It calls Close() at the end of the block.
type blockBuilder interface {
	// EndStmt ends a statement
	EndStmt()

	// AddBlock appends a new block entry
	AddBlock() blockBuilder

	// AddToken appends a new token entry
	AddTok(t *Tok)

	// Close closes this block
	Close()
}
