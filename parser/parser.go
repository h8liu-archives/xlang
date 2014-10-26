package parser

import (
	"fmt"
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
}

func (p *parser) parse(lex *Lexer) {
	for lex.Scan() {
		// TODO: process the token
		lex.Token()
	}
}

// Parse parses a file from the input stream.
func Parse(file string, r io.ReadCloser) (Block, ErrList) {
	lex := Lex(file, r)
	p := new(parser)
	p.parse(lex)

	ioErr := lex.IOErr()
	if ioErr != nil {
		return nil, singleErr(ioErr)
	}

	lexErrs := lex.Errors()
	if lexErrs.Len() > 0 {
		return nil, lexErrs
	}

	// TODO: lexing passed, now do the parsing
	return nil, singleErr(fmt.Errorf("parser not implemented"))
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
