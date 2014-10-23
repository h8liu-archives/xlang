package lexer

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Lexer struct {
	file string
	r    io.ReadCloser
}

func open(path string) (*Lexer, error) {
	f, e := os.Open(path)
	if e != nil {
		return nil, e
	}

	return newLexer(path, f), nil
}

func lex(file, s string) *Lexer {
	r := ioutil.NopCloser(strings.NewReader(s))
	return newLexer(file, r)
}

func newLexer(file string, r io.ReadCloser) *Lexer {
	ret := new(Lexer)
	ret.file = file
	ret.r = r

	return ret
}
