package main

import (
	"fmt"

	"github.com/h8liu/xlang/parser"
)

const prog = `
var x = 3a
var y = 4
var z = x + y

func main() {
	print(x)
}

print(z) // some comment
/* okay */
`

func main() {
	lex := parser.LexString("test.x", prog)

	for lex.Scan() {
		t := lex.Token()
		fmt.Println(t)
	}
}
