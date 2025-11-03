package errors

type Op string

type Error interface {
	error
}

type DetailedError struct {
	op    Op
	cause error
	msg   string // Human-readable error message
}

// New creates a new error with the given Op and a default message: "Internal system error."
// The Cause is set to nil.
func New(op Op) *DetailedError {
	return &DetailedError{
		op:    op,
		cause: nil,
		msg:   "Internal system error.",
	}
}

func (e *DetailedError) Error() string {
	return e.msg
}

func (e *DetailedError) Msg(msg string) {
	e.msg = msg
}
