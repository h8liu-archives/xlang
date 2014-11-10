package main

import (
	"fmt"

	"github.com/h8liu/xlang/e8/img"
	"github.com/h8liu/xlang/e8/mem"
)

func mainRun(args []string) {
	if len(args) != 1 {
		panic("need exactly one arg")
	}

	f := args[0]
	vm, e := img.Open(f)
	if e != nil {
		printError(e)
		return
	}

	vm.SetPC(mem.SegCode)
	n := 0
	for !vm.Halted() {
		n += vm.Run(1000)
	}

	if !vm.RIP() {
		fmt.Printf("(ret=%d)\n", vm.HaltValue())
		if vm.AddrError() {
			printError(fmt.Errorf("vm halted on address error"))
		}
	}

	fmt.Printf("(%d cycles)\n", n)
}
