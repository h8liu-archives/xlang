package parser

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type parser struct {
	lex   *Lexer
	block Block
	errs  *ErrList

	eofErrored bool
}

func newParser(lex *Lexer) *parser {
	ret := new(parser)
	ret.lex = lex
	ret.errs = NewErrList()

	return ret
}

func endStmtToken(t *Tok) bool {
	return t.Type == TypeOperator &&
		(t.Lit == "\n" || t.Lit == ";" || t.Lit == "}")
}

func endBlockToken(t *Tok) bool {
	return t.Type == TypeOperator && t.Lit == "}"
}

func startBlockToken(t *Tok) bool {
	return t.Type == TypeOperator && t.Lit == "{"
}

func (p *parser) parseEntry() *Entry {
	var t *Tok

	for {
		if !p.lex.Scan() {
			return nil
		}

		t = p.lex.Token()
		if t.Type != TypeComment {
			break
		}
	}

	if endStmtToken(t) {
		return nil
	}

	if startBlockToken(t) {
		b := p.parseBlock()
		b.Lbrace = t

		if p.lex.EOF() && !p.eofErrored {
			p.errs.Log(p.lex.Pos(), "unexpected EOF")
			p.eofErrored = true
		} else {
			b.Rbrace = p.lex.Token()
		}

		return &Entry{Block: b}
	}

	return &Entry{Tok: t}
}

// parseStmt returns a statement, or nil when reaching end of a block
func (p *parser) parseStmt() *Stmt {
	ret := new(Stmt)

	for {
		e := p.parseEntry()
		if e == nil {
			ret.End = p.lex.Token()
			break
		}
		ret.Entries = append(ret.Entries, e)
	}

	if ret.Entries == nil {
		if p.lex.EOF() || endBlockToken(p.lex.Token()) {
			return nil
		}
		return ret
	}

	return ret
}

func (p *parser) parseBlock() *Block {
	var b = new(Block)
	for {
		s := p.parseStmt()
		if s == nil {
			break
		}
		b.Stmts = append(b.Stmts, s)

		if p.lex.EOF() || endBlockToken(p.lex.Token()) {
			break
		}
	}
	return b
}

func (p *parser) parse() *Block {
	ret := p.parseBlock()
	if !p.lex.EOF() {
		t := p.lex.Token()
		if !endBlockToken(t) {
			panic("bug")
		}

		p.errs.Log(t.Pos, "unmatched }")

		for p.lex.Scan() {
			// eat up the rest of the tokens
		}
	}
	return ret
}

// Parse parses a file from the input stream.
func Parse(file string, r io.ReadCloser) (*Block, *ErrList) {
	lex := Lex(file, r)
	p := newParser(lex)
	ret := p.parse()

	ioErr := lex.IOErr()
	if ioErr != nil {
		return nil, singleErr(ioErr)
	}

	lexErrs := lex.Errors()
	if lexErrs.Len() > 0 {
		return nil, lexErrs
	}

	if p.errs.Len() > 0 {
		return nil, p.errs
	}

	return ret, nil
}

// ParseFile parses a file (on the file system).
func ParseFile(path string) (*Block, *ErrList) {
	f, e := os.Open(path)
	if e != nil {
		return nil, singleErr(e)
	}

	return Parse(path, f)
}

// ParseStr parse a file from a string.
func ParseStr(file, s string) (*Block, *ErrList) {
	r := ioutil.NopCloser(strings.NewReader(s))
	return Parse(file, r)
}
