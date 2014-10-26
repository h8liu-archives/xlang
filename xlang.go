package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/h8liu/xlang/parser"
)

const prog = `var x = 3a
var y = 4
var z = x + y

func main() {
	print(x)
}

print(z) // some comment
/* okay */
`

var (
	runWeb = flag.Bool("-web", false, "run web server")
)

func main() {
	if *runWeb {
		webMain()
		return
	}

	_, errs := parser.ParseStr("test.x", prog)

	if errs != nil {
		for errs.Scan() {
			fmt.Println(errs.Error())
		}
	}
}

func webMain() {
	server := http.FileServer(http.Dir("."))
	http.Handle("/", server)

	for {
		e := http.ListenAndServe(":8000", nil)
		if e != nil {
			log.Fatal(e)
		}
	}
}
