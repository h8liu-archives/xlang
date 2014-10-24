package parser

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// lexer is a token scanner
type lexer struct {
	e    error
	c    *cursor
	last *Tok
	hold *Tok
}

func newLexer(file string, r io.ReadCloser) *lexer {
	ret := new(lexer)
	ret.c = newCursor(file, r)

	return ret
}

func (lex *lexer) emitTok(t *Tok) {
	lex.hold = t
	if lex.hold.Type != TypeComment {
		lex.last = t
	}
}

func (lex *lexer) emitType(t Type) {
	tok := lex.c.Token(t)
	lex.emitTok(tok)
}

func (lex *lexer) skipEndl() bool {
	if lex.last == nil {
		return true
	}
	t := lex.last.Type

	switch t {
	case TypeIdent, TypeInt, TypeFloat, TypeString:
		return false
	case TypeOperator:
		lit := lex.last.Lit
		return !(lit == "}" || lit == "]" || lit == ")")
	}

	return true
}

func (lex *lexer) scanInt() {
	for lex.c.Scan() {
		r := lex.c.Next()
		if !isDigit(r) {
			break
		}
	}

	lex.emitType(TypeInt)
}

func (lex *lexer) scanIdent() {
	for lex.c.Scan() {
		r := lex.c.Next()
		if !isDigit(r) && !isLetter(r) {
			break
		}
	}

	lex.emitType(TypeIdent)
}

func (lex *lexer) scanLineComment() {
	for lex.c.Scan() {
		r := lex.c.Next()
		if r == '\n' {
			break
		}
	}

	lex.emitType(TypeComment)
}

func (lex *lexer) scanBlockComment() {
	var star bool
	var complete bool

	for lex.c.Scan() {
		r := lex.c.Next()
		if star && r == '/' {
			lex.c.Accept()
			complete = true
			break
		}

		star = r == '*'
	}

	lex.emitType(TypeComment)
	if !complete {
		// TODO: report imcomplete block comment parsing error
	}
}

func (lex *lexer) isWhite(r rune) bool {
	if isWhite(r) {
		return true
	}

	if r == '\n' {
		return lex.skipEndl()
	}

	return false
}

func (lex *lexer) skipWhite() {
	if lex.c.EOF() || !lex.isWhite(lex.c.Next()) {
		return // no white to skip
	}

	for lex.c.Scan() {
		r := lex.c.Next()
		if !lex.isWhite(r) {
			break
		}
	}

	lex.c.Discard()
}

func (lex *lexer) scanOperator() {
	r := lex.c.Next()

	if lex.c.Scan() && r == '/' {
		r2 := lex.c.Next()

		if r2 == '/' {
			lex.scanLineComment()
			return
		} else if r2 == '*' {
			lex.scanBlockComment()
			return
		}
	}

	if r == '\n' && lex.skipEndl() {
		panic("bug")
	}

	lex.emitType(TypeOperator)
}

func (lex *lexer) scanInvalid() {
	lex.c.Accept()
	lex.emitType(TypeInvalid)
}

func (lex *lexer) Scan() bool {
	lex.skipWhite()

	if lex.c.EOF() {
		return false
	}
	r := lex.c.Next()

	if isDigit(r) {
		lex.scanInt()
	} else if isLetter(r) {
		lex.scanIdent()
	} else if isOperator(r) {
		lex.scanOperator()
	} else {
		lex.scanInvalid()
	}

	return true
}

func (lex *lexer) Err() error {
	return lex.e
}

func lexOpen(path string) (*lexer, error) {
	f, e := os.Open(path)
	if e != nil {
		return nil, e
	}

	return newLexer(path, f), nil
}

func lexString(file, s string) *lexer {
	r := ioutil.NopCloser(strings.NewReader(s))
	return newLexer(file, r)
}
