package parser

import (
	"bytes"
)

type cursor struct {
	file     string
	s        *runeScanner
	buf      *bytes.Buffer
	row, col int

	head rune
	eof  bool
}

func newCursor(f string, s *runeScanner) *cursor {
	ret := new(cursor)
	ret.s = s
	ret.file = f
	ret.buf = new(bytes.Buffer)

	// bootstrap the header
	if !ret.s.Scan() {
		ret.eof = true
	}

	return ret
}

// Scan returns true if any progress is made.
func (c *cursor) Scan() bool {
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
	ret.Pos = &Pos{
		File: c.file,
		Row:  c.row,
		Col:  c.col,
	}

	// reset the saving buffer
	c.buf.Reset()
	c.row, c.col = c.s.Pos()

	return ret
}

// EOF checkes if the head cursor is pointing to the end of the file.
func (c *cursor) EOF() bool {
	return c.eof
}

// Peek returns the rune pointed by the head cursor.
// If the head cursor is pointing to the end of the file, it will panic.
func (c *cursor) Peek() rune {
	if c.eof {
		panic("cursor pointing to the end of the file")
	}
	return c.s.Rune()
}

// Err returns any error encountered using the scanner
func (c *cursor) Err() error {
	return c.s.Err()
}
