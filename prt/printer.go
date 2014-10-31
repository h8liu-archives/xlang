// Package prt provides a general indent printing interface
// and an implementation.
package prt

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Printer provides indented line printing
type Printer struct {
	Prefix string
	Indent string
	Shift  int
	Writer io.Writer
	Error  error
}

var _ Iface = new(Printer)

// New creates a new printer that writes to w.
// If w is nil, all prints to the printer will be noops.
func New(w io.Writer) *Printer {
	if w == nil {
		w = noop
	}

	return &Printer{
		Indent: "\t",
		Writer: w,
	}
}

// Stdout creates a new printer that prints to os.Stdout
func Stdout() *Printer {
	return New(os.Stdout)
}

func (p *Printer) p(n *int, a ...interface{}) {
	if p.Error != nil {
		return
	}

	i, e := fmt.Fprint(p.Writer, a...)
	p.Error = e
	*n += i
}

func (p *Printer) pln(n *int, a ...interface{}) {
	if p.Error != nil {
		return
	}

	i, e := fmt.Fprintln(p.Writer, a...)
	p.Error = e
	*n += i
}

func (p *Printer) pf(n *int, format string, a ...interface{}) {
	if p.Error != nil {
		return
	}

	i, e := fmt.Fprintf(p.Writer, format, a...)
	p.Error = e
	*n += i
}

func (p *Printer) pre(n *int) {
	p.p(n, p.Prefix)
	for i := 0; i < p.Shift; i++ {
		p.p(n, p.Indent)
	}
}

// Print prints a line of args using fmt.Print
func (p *Printer) Print(a ...interface{}) (int, error) {
	n := 0
	p.pre(&n)
	p.p(&n, a...)
	p.pln(&n)

	return n, p.Error
}

// Println prints a line of args using fmt.Println
func (p *Printer) Println(a ...interface{}) (int, error) {
	n := 0
	p.pre(&n)
	p.pln(&n, a...)

	return n, p.Error
}

// Printf prints a line of formatted args using fmt.Printf
func (p *Printer) Printf(format string, a ...interface{}) (int, error) {
	n := 0
	p.pre(&n)
	p.pf(&n, format, a...)
	p.pln(&n)

	return n, p.Error
}

// ShiftIn adds one level of indentation
func (p *Printer) ShiftIn() {
	p.Shift++
}

// ShiftOut removes one level of indentation
func (p *Printer) ShiftOut(a ...interface{}) {
	if p.Shift == 0 {
		panic("shift already left most")
	}
	p.Shift--

	if len(a) > 0 {
		p.Print(a...)
	}
}

// String creates a string of a Printable object.
func String(p Printable) string {
	buf := new(bytes.Buffer)
	dev := New(buf)
	p.PrintTo(dev)
	return buf.String()
}

// Print prints a printable object to os.Stdout
func Print(p Printable) {
	p.PrintTo(Stdout())
}
