// e8 command loads and simulates a E8 virtual machine image.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		panic("todo: intro")
		return
	}

	cmd := args[1]
	subArgs := args[2:]

	switch cmd {
	case "dasm":
		mainDasm(subArgs)
	case "run":
		mainRun(subArgs)
	default:
		if strings.HasSuffix(cmd, ".e8") {
			mainRun(args[1:])
		} else {
			fmt.Fprintf(os.Stderr, "e8: unknown subcommand %q.\n", cmd)
			fmt.Fprintf(os.Stderr, "Run 'e8 help' for usage.\n")
		}
	}
}

func printError(e error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", e)
}
