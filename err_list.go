package xlang

import (
	"fmt"
)

const maxListError = 20

type ErrList struct {
	errs    []*Error
	scanned bool
	scanPtr int
	hold    *Error
}

func NewErrList() *ErrList {
	ret := new(ErrList)

	return ret
}

// Log appends an error to the error list.
// It panics when error is nil.
func (lst *ErrList) Log(p *Pos, f string, args ...interface{}) {
	if len(lst.errs) < maxListError {
		e := &Error{p, fmt.Sprintf(f, args...)}
		lst.errs = append(lst.errs, e)
	}
}

// Len returns the number of errors in the list.
func (lst *ErrList) Len() int {
	return len(lst.errs)
}

// Scan returns if there is an error for scanning.
func (lst *ErrList) Scan() bool {
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
func (lst *ErrList) Error() *Error {
	if lst.hold == nil {
		panic("invalid operation")
	}

	return lst.hold
}
