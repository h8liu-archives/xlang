package main

import (
	"fmt"

	"github.com/h8liu/xlang"
)

var ident = 0

func printIdent() {
	for i := 0; i < ident; i++ {
		fmt.Print("    ")
	}
}

func printBlock(b xlang.Block) {
	fmt.Println("{")
	ident++

	for _, stmt := range b {
		printStmt(stmt)
	}

	ident--
	printIdent()
	fmt.Print("}")
}

func printStmt(s xlang.Stmt) {
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
	block, errs := xlang.ParseStr("test.x", prog)
	if errs != nil {
		for errs.Scan() {
			fmt.Println(errs.Error())
		}
	} else {
		printBlock(block)
		fmt.Println()
	}
}
