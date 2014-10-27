package parser

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func singleErr(e error) ErrList {
	errs := newErrList()
	errs.Log(nil, e.Error())
	return errs
}

type parser struct {
	lex   *Lexer
	block Block
	errs  *errList

	eofErrored bool
}

func newParser(lex *Lexer) *parser {
	ret := new(parser)
	ret.lex = lex
	ret.errs = newErrList()

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
	return t.Type == TypeOperator && t.Lit == "}"
}

func (p *parser) parseEntry() *Entry {
	if p.lex.Scan() {
		t := p.lex.Token()
		if endStmtToken(t) {
			return nil
		}
		if startBlockToken(t) {
			b := p.parseBlock()

			if p.lex.EOF() && !p.eofErrored {
				p.errs.Log(p.lex.Pos(), "unexpected EOF")
				p.eofErrored = true
			}

			return &Entry{Block: b}
		}
		return &Entry{Tok: t}
	}
	return nil
}

// parseStmt returns a statement, or nil when reaching end of a block
func (p *parser) parseStmt() Stmt {
	var ret Stmt

	for {
		e := p.parseEntry()
		if e == nil {
			break
		}
		ret = append(ret, e)
	}

	if ret == nil {
		if p.lex.EOF() || endBlockToken(p.lex.Token()) {
			return nil
		} else {
			return make(Stmt, 0)
		}
	}

	return ret
}

func (p *parser) parseBlock() Block {
	var b Block
	for {
		s := p.parseStmt()
		if s == nil {
			break
		}
		b = append(b, s)
	}
	return b
}

func (p *parser) parse(lex *Lexer) Block {
	return p.parseBlock()
}

// Parse parses a file from the input stream.
func Parse(file string, r io.ReadCloser) (Block, ErrList) {
	lex := Lex(file, r)
	p := new(parser)
	ret := p.parse(lex)

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
func ParseFile(path string) (Block, ErrList) {
	f, e := os.Open(path)
	if e != nil {
		return nil, singleErr(e)
	}

	return Parse(path, f)
}

// ParseStr parse a file from a string.
func ParseStr(file, s string) (Block, ErrList) {
	r := ioutil.NopCloser(strings.NewReader(s))
	return Parse(file, r)
}
