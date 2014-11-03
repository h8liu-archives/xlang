package xc_test

import (
	"bytes"
	"strings"

	"github.com/h8liu/xlang/xc"

	"testing"
)

func TestFunc(t *testing.T) {
	o := func(s string, expect string) {
		src := xc.NewStrSource("text.x", s)
		obj, errs := src.CompileFunc()
		if errs != nil {
			t.Errorf("error on parsing: %q", s)
			return
		}

		buf := new(bytes.Buffer)
		obj.Sim(buf)
		out := strings.TrimSpace(buf.String())

		if out != expect {
			t.Errorf("running %q, got %q, expect %q",
				s, out, expect,
			)
		}
	}

	e := func(s string) {
		src := xc.NewStrSource("test.x", s)
		_, errs := src.CompileFunc()
		if errs == nil {
			t.Errorf("expect error on parsing: %q", s)
		}
	}

	o("", "")
	o(";;;", "")
	o("print(3)", "3")
	o("print(33);", "33")
	o("var x=3; print(x)", "3")
	o("var x=3; var y=4; print(x+y)", "7")
	o("var x; x=4; print(x)", "4")
	o("var x=4; print(x-3)", "1")
	o("var x; var y; x,y = 3,4; print(x-y)", "-1")
	o("var x,y; x,y = 3,4; print(x-y)", "-1")
	o("var x,y = 3,4; print(x-y)", "-1")
	o("var x,y; y=4; print(x-y)", "-4")

	e("print(a)")
	e("var y = x-3")
	e("var x,y = 3")
}
