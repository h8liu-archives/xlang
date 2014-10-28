package xlang

import (
	"bufio"
	"io"
)

type runeScanner struct {
	rc       io.ReadCloser
	r        *bufio.Reader
	row, col int
	closed   bool
	scanned  bool

	hold rune
	e    error
}

func newRuneScanner(rc io.ReadCloser) *runeScanner {
	ret := new(runeScanner)
	ret.r = bufio.NewReader(rc)
	ret.rc = rc

	return ret
}

// Pos returns the position of the current rune.
// If the scanner is closed (reaching EOF or got an error), Pos returns
// the position of the last rune read.
func (s *runeScanner) Pos() (row, col int) {
	return s.row + 1, s.col + 1
}

// Scan moves forward the rune cursor by one.
// It returns false if it reaches EOF or encounters an error.
func (s *runeScanner) Scan() bool {
	if s.closed {
		return false
	}

	wasEndl := s.hold == '\n'

	s.hold, _, s.e = s.r.ReadRune()
	if s.e != nil {
		e := s.rc.Close()
		if s.e == io.EOF {
			s.e = e // replace EOF with closing error
		}

		s.closed = true
		return false
	}

	if wasEndl {
		s.row++
		s.col = 0
	} else if s.scanned {
		s.col++
	} else {
		s.scanned = true
	}

	return true
}

// Err returns the first error encountered on scanning.
// Returns nil if no error.
func (s *runeScanner) Err() error {
	return s.e
}

// Rune returns the rune pointed by the cursor.
// It panics when the scanner is closed already or when an error was met.
// Please use only call Rune() when Scan() returns true.
// Could be called multiple times.
func (s *runeScanner) Rune() rune {
	if s.closed {
		panic("scanner closed")
	}
	if s.e != nil {
		panic("scanning error encountered")
	}
	if !s.scanned {
		panic("not scanned yet")
	}
	return s.hold
}
