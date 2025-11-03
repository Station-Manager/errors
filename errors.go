package errors

type Op string

type Error interface {
	error
}

type DetailedError struct {
	Op    Op
	Cause error
	Msg   string // Human-readable error message
}

func (e *DetailedError) Error() string {
	return e.Msg
}

func New(op Op) *DetailedError {
	return &DetailedError{
		Op:    op,
		Cause: nil,
		Msg:   "Internal system error.",
	}
}
