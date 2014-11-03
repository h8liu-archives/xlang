package main

import (
	"fmt"

	"github.com/h8liu/xlang/xc"
)

var ident = 0

func main() {
	src := xc.NewStrSource("test.x", prog)
	obj, errs := src.CompileFunc()
	if errs != nil {
		for errs.Scan() {
			fmt.Println(errs.Error())
		}
		return
	}

	obj.Sim()
}
