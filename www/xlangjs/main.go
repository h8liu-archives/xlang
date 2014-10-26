package main

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/gopherjs/gopherjs/js"

	"github.com/h8liu/xlang/parser"
)

var code = `
var x = 3
var y = 4

func main() {
	println(x + y)
}
`

func main() {
	js.Global.Set("xlang", map[string]interface{}{
		"parseTokens": parseTokens,
		"parse":       parse,
	})
}

func _parse(file, code string) (block, errs string) {
	_, es := parser.ParseStr(file, code)
	if es != nil {
		buf := new(bytes.Buffer)

		for es.Scan() {
			fmt.Fprintf(buf, `<div class="error">%s</div>`,
				template.HTMLEscapeString(es.Error().String()),
			)
		}

		return "", buf.String()
	}

	return "", ""
}

func parse(file, code string) map[string]interface{} {
	ret := make(map[string]interface{})
	block, errs := _parse(file, code)
	ret["block"] = block
	ret["errs"] = errs

	return ret
}

func parseTokens(file, code string) string {
	lex := parser.LexStr(file, code)
	out := new(bytes.Buffer)

	lines := strings.Split(code, "\n")
	toks := make(map[uint64]*parser.Tok)
	for lex.Scan() {
		t := lex.Token()
		toks[(uint64(t.Row)<<32)+uint64(t.Col)] = t
	}

	var curTok *parser.Tok
	var curPos int
	var curLit []rune

	emit := func(row, col int, r rune) {
		index := (uint64(row) << 32) + uint64(col)
		if curTok == nil {
			tok := toks[index]
			if tok != nil {
				curTok = tok
				class := "na"

				switch tok.Type {
				case parser.TypeIdent:
					class = "ident"
				case parser.TypeInt, parser.TypeFloat:
					class = "num"
				case parser.TypeOperator:
					class = "operator"
				case parser.TypeKeyword:
					class = "keyword"
				case parser.TypeComment:
					class = "comment"
				}

				if tok.Lit == "\n" && tok.Type == parser.TypeOperator {
					class = "implicit-semi"
				}

				fmt.Fprintf(out, `<span class="%s">`, class)

				if class == "implicit-semi" {
					fmt.Fprintf(out, ";")
				}

				curPos = 0
				curLit = nil
				for _, r := range tok.Lit {
					curLit = append(curLit, r)
				}
			}
		} else {
			if toks[index] != nil {
				fmt.Println("overlapping token %d:%d\n", row, col)
			}
		}

		if r == '\t' {
			fmt.Fprint(out, "&nbsp;&nbsp;&nbsp;&nbsp;")
		} else if r == '\n' {
			fmt.Fprint(out, "<br/>")
		} else if r == ' ' {
			fmt.Fprint(out, "&nbsp;")
		} else {
			fmt.Fprint(out, template.HTMLEscapeString(string(r)))
		}

		if curTok != nil {
			if curLit[curPos] != r {
				fmt.Printf("mismatch at %d:%d\n", row, col)
			}

			curPos++
			if curPos >= len(curTok.Lit) {
				curTok = nil
				fmt.Fprintf(out, "</span>")
			}
		}
	}

	for row, line := range lines {
		for col, r := range line {
			emit(row+1, col+1, r)
		}
		emit(row+1, len(line)+1, '\n')
	}

	return out.String()
}
