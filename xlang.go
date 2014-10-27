package main

import (
	"fmt"

	"github.com/h8liu/xlang/parser"
)

const prog = `var x = 3a
var y = 4
var z = x + y

func main() {
	{
		print(x);
	}
	;
}

print(z) // some comment
`

var ident = 0

func printIdent() {
	for i := 0; i < ident; i++ {
		fmt.Print("    ")
	}
}

func printBlock(b parser.Block) {
	fmt.Println("{")
	ident++
	
	for _, stmt := range b {
		printStmt(stmt)
	}

	ident--
	printIdent()
	fmt.Print("}")
}

func printStmt(s parser.Stmt) {
	printIdent()

	if len(s) == 0 {
		fmt.Printf("<empty>")
	} else {
		for i, e := range s {
			if i > 0 {
				fmt.Printf(" ")
			}

			if e.Block != nil {
				printBlock(e.Block)
			} else {
				t := e.Tok
				fmt.Printf("%s:%q", t.Type.ShortStr(), t.Lit)
			}
		}
	}

	fmt.Println()
}

func main() {
	block, errs := parser.ParseStr("test.x", prog)
	if errs != nil {
		for errs.Scan() {
			fmt.Println(errs.Error())
		}
	}
	
	printBlock(block) 
	fmt.Println()
}
