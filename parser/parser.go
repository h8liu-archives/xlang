package parser

import (
	"fmt"
	"io"
)

// Parsed contains a parsed file.
type Parsed struct {
	Errors ErrList // the list of parsing errors
	Block  Block   // the returned block
}

func parsedErrs(errs ErrList) *Parsed {
	return &Parsed{Errors: errs}
}

func parsedErr(e error) *Parsed {
	errs := newErrList()
	errs.Log(nil, e.Error())
	return &Parsed{Errors: errs}
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
func Parse(file string, r io.ReadCloser) *Parsed {
	lex := Lex(file, r)
	p := new(parser)
	p.parse(lex)

	ioErr := lex.IOErr()
	if ioErr != nil {
		return parsedErr(ioErr)
	}

	lexErrs := lex.Errors()
	if lexErrs.Len() > 0 {
		return parsedErrs(lexErrs)
	}

	// TODO: lexing passed, now do the parsing
	return parsedErr(fmt.Errorf("parser not implemented"))
}
