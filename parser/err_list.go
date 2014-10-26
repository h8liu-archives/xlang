package parser

import (
	"fmt"
)

// ErrList defines the interface for iterating over
// an error list.
type ErrList interface {
	// Len returns the total count of errors in the list
	Len() int

	// Scan returns if there is an error for scanning.
	Scan() bool

	// Error returns the error.
	Error() *Error
}

type errList struct {
	maxError int // maximum number of errors to record in the list
	errs     []*Error

	scanned bool
	scanPtr int
	hold    *Error
}

func newErrList() *errList {
	ret := new(errList)
	ret.maxError = 20

	return ret
}

// Log appends an error to the error list.
// It panics when error is nil.
func (lst *errList) Log(p *Pos, f string, args ...interface{}) {
	if len(lst.errs) < lst.maxError {
		e := &Error{p, fmt.Sprintf(f, args...)}
		lst.errs = append(lst.errs, e)
	}
}

// Len returns the number of errors in the list.
func (lst *errList) Len() int {
	return len(lst.errs)
}

// Scan returns if there is an error for scanning.
func (lst *errList) Scan() bool {
	if lst.scanned {
		lst.scanPtr++
	} else {
		lst.scanned = true
	}

	if lst.scanPtr < len(lst.errs) {
		lst.hold = lst.errs[lst.scanPtr]
		return true
	}

	lst.hold = nil
	return false
}

// Error returns the current scaned error.
// It returns nil for invalid operations.
func (lst *errList) Error() *Error {
	if lst.hold == nil {
		panic("invalid operation")
	}

	return lst.hold
}
