package errors

import (
	stderr "errors"
)

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

// AsDetailedError attempts to cast the given error to a DetailedError instance.
func AsDetailedError(err error) (*DetailedError, bool) {
	var dErr *DetailedError
	if stderr.As(err, &dErr) {
		return dErr, true
	}
	return nil, false
}

// Error implements the error interface.
func (e *DetailedError) Error() string {
	if e == nil {
		return ""
	}
	return e.msg
}

// Msg sets the human-readable error message for a DetailedError instance.
func (e *DetailedError) Msg(msg string) *DetailedError {
	if e == nil {
		return nil
	}
	e.msg = msg
	return e
}

// Err sets the cause for a DetailedError instance.
func (e *DetailedError) Err(err error) *DetailedError {
	if e == nil {
		return nil
	}
	e.cause = err
	return e
}

// Cause returns the cause of a DetailedError instance.
func (e *DetailedError) Cause() error {
	if e == nil {
		return nil
	}
	if e.cause == nil {
		return nil
	}
	return e.cause
}
