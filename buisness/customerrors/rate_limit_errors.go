package customerrors

import "errors"

// RateLimitError is used to pass an error during the request through the
// application with web specific context.
type RateLimitError struct {
	Err    error
	Status int
}

// NewRateLimitError wraps a provided error with an HTTP status code. This
// function should be used when handlers encounter expected errors.
func NewRateLimitError(err error, status int) error {
	return &RateLimitError{err, status}
}

// Error implements the error interface. It uses the default message of the
// wrapped error. This is what will be shown in the services' logs.
func (re *RateLimitError) Error() string {
	return re.Err.Error()
}

// IsRateLimitError checks if an error of type RequestError exists.
func IsRateLimitError(err error) bool {
	var re *RateLimitError
	return errors.As(err, &re)
}

// GetRequestError returns a copy of the RequestError pointer.
func GetRateLimitError(err error) *RateLimitError {
	var re *RateLimitError
	if !errors.As(err, &re) {
		return nil
	}
	return re
}
