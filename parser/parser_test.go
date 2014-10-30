package parser_test

import (
	"github.com/h8liu/xlang/parser"

	"testing"
)

func TestParser(t *testing.T) {
	o := func(s string) {
		_, es := parser.ParseStr("test.x", s)
		if es != nil {
			t.Errorf("error on parsing: %q", s)
		}
	}

	e := func(s string) {
		_, es := parser.ParseStr("test.x", s)
		if es == nil {
			t.Errorf("expect error on parsing: %q", s)
		}
	}

	o("")
	o("\n")
	o("{}")
	o("a b c")
	o("{{n}}")
	o("{{;}}")
	o(";;;;;")
	o("{;{;};}")
	o("func main() { return }")

	e("{")
	e("}")
	e("{{}")
	e("{}}")
	e("{};}")
	e("{};\n}")
	e("/* incomplete block comment")
}
