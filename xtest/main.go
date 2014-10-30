package main

import (
	"fmt"

	"github.com/h8liu/xlang/parser"
	"github.com/h8liu/xlang/xc"
)

var ident = 0

func printIdent() {
	for i := 0; i < ident; i++ {
		fmt.Print("    ")
	}
}

func printBlock(b *parser.Block) {
	fmt.Println("{")
	ident++

	for _, stmt := range b.Stmts {
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
	src := xc.NewStrSource("test.xpr", prog)
	b, errs := src.BuildExprsAST()
	if errs != nil {
		for errs.Scan() {
			fmt.Println(errs.Error())
		}
	} else {
		for _, n := range b.Nodes {
			fmt.Println(xc.ExprStr(n))
		}
	}
}
