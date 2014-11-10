package vm

import (
	"io"

	"e8vm.net/e8/mem"
)

// SysPage is a special system page that is always mapped to address 0.
type SysPage struct {
	AddrError   bool  // if anything read or written to address 0-3
	Halt        bool  // if halt byte (address 4) is written
	HaltValue   uint8 // the halt value written on address 4
	StdoutError error // last IO error on flushing stdout

	stdin  chan byte
	stdout chan byte
}

// Special addresses in a system page.
const (
	// Halt is the address for halting
	// write: sets the halt value, halts the machine
	Halt = 8

	// Stdout is the address for output a char
	// read: if stdout is ready for output
	// write: output a byte to stdout
	Stdout = 9

	// StdinReady is the address for testing if an input is ready.
	// read: if stdin is ready
	StdinReady = 10

	// Stdin is the address for fetching the input.
	// read: fetch a byte if any, returns 0 if stdin is not ready
	Stdin = 11
)

var _ mem.Page = new(SysPage)

// NewSysPage creates a system page
func NewSysPage() *SysPage {
	ret := new(SysPage)
	ret.stdin = make(chan byte, 32)
	ret.stdout = make(chan byte, 32)

	return ret
}

// Halted returns if the state is halted
func (p *SysPage) Halted() bool {
	if p == nil {
		return false
	}

	return p.Halt
}

// Reset clears errors; the system is no longer halting afterwards.
func (p *SysPage) Reset() {
	if p == nil {
		return
	}
	p.AddrError = false
	p.Halt = false
}

func (p *SysPage) addrError() {
	p.AddrError = true
	p.Halt = true
	p.HaltValue = 0xff
}

// Read reads a byte at address offset
func (p *SysPage) Read(offset uint32) uint8 {
	if offset < 8 {
		p.addrError()
		return 0
	}

	switch offset {
	case Stdout: // stdout ready
		if len(p.stdout) < cap(p.stdout) {
			return 0 // ready
		}
		return 1 // busy
	case StdinReady: // stdin ready
		if len(p.stdin) > 0 {
			return 0
		}
		return 1 // invalid
	case Stdin: // stdin value
		if len(p.stdin) > 0 {
			return <-p.stdin
		}
		return 0
	default:
		return 0
	}
}

// Write writes a byte at address offset
func (p *SysPage) Write(offset uint32, b uint8) {
	if offset < 8 {
		p.addrError()
		return
	}

	switch offset {
	case Halt: // halt
		p.Halt = true
		p.HaltValue = b
	case Stdout: // stdout
		if len(p.stdout) < cap(p.stdout) {
			p.stdout <- b
		}
	}
}

// FlushStdout flushes buffered stdout bytes to Writer.
// Errors will be stored on p.IoError
func (p *SysPage) FlushStdout(w io.Writer) {
	if p == nil {
		return
	}

	for len(p.stdout) > 0 {
		b := <-p.stdout
		_, e := w.Write([]byte{b})
		if e != nil {
			p.StdoutError = e
		}
	}
}
