package main

import (
	"bytes"
	"fmt"

	"github.com/h8liu/xlang/parser"
)

func main() {

}

func parseTokens(file, code string) string {
	lex := parser.LexString(file, code)
	out := new(bytes.Buffer)

	for lex.Scan() {
		t := lex.Token()
		fmt.Fprintln(out, t)
	}
	return out.String()
}
