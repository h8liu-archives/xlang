package parser

import (
	"bytes"
	"io"
)

type cursor struct {
	file     string
	s        *runeScanner
	buf      *bytes.Buffer
	row, col int

	head rune
	eof  bool
}

func newCursor(f string, r io.ReadCloser) *cursor {
	ret := new(cursor)
	ret.s = newRuneScanner(r)
	ret.file = f
	ret.buf = new(bytes.Buffer)

	ret.row, ret.col = ret.s.Pos()

	// bootstrap the header
	if !ret.s.Scan() {
		ret.eof = true
	}

	return ret
}

// Scan returns true if any progress is made.
// When a progress is made,
func (c *cursor) Accept() bool {
	if c.eof {
		// pointing to EOF already, no progress made
		return false
	}

	c.buf.WriteRune(c.s.Rune())
	if !c.s.Scan() {
		c.eof = true // will return false next time
	}

	return true
}

// ScanNext returns true if progress is made and the pointer
// is not pointing to EOF.
// When it returns true, it is always safe to call Next().
func (c *cursor) Scan() bool {
	return c.Accept() && !c.EOF()
}

// Pos returns the current position of the head cursor.
func (c *cursor) Pos() *Pos {
	return &Pos{
		File: c.file,
		Row:  c.row,
		Col:  c.col,
	}
}

// Token wraps the the runes in the buffer into a token.
// It then resets the buffer cursor.
// It panics when nothing is buffered.
func (c *cursor) Token(t Type) *Tok {
	if c.buf.Len() == 0 {
		panic("nothing buffered")
	}

	ret := new(Tok)
	ret.Type = t
	ret.Lit = c.buf.String()
	ret.Pos = c.Pos()

	c.resetBuf()

	return ret
}

func (c *cursor) Discard() {
	c.resetBuf()
}

func (c *cursor) resetBuf() {
	c.buf.Reset()
	c.row, c.col = c.s.Pos()
}

func (c *cursor) Buffered() string {
	return c.buf.String()
}

// EOF checkes if the head cursor is pointing to the end of the file.
func (c *cursor) EOF() bool {
	return c.eof
}

// Next returns the rune pointed by the head cursor.
// If the head cursor is pointing to the end of the file, it will panic.
func (c *cursor) Next() rune {
	if c.eof {
		panic("cursor pointing to the end of the file")
	}
	return c.s.Rune()
}

// Err returns any error encountered using the scanner
func (c *cursor) Err() error {
	return c.s.Err()
}
