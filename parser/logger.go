package parser

type ErrLogger interface {
	Log(e *Error)
}
