package util

import (
	"fmt"
	"strings"
)

// ErrorHandler helps collect and manage multiple errors
type ErrorHandler struct {
	errs []error
}

// NewErrorHandler creates a new ErrorHandler
func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{
		errs: make([]error, 0),
	}
}

// Push adds an error to the handler if it's not nil
// If field is provided, it will be included in the error message
func (h *ErrorHandler) Push(field string, err error) {
	if err != nil {
		errMsg := err.Error()
		if field != "" {
			errMsg = fmt.Sprintf("%s: %s", field, errMsg)
		}
		h.errs = append(h.errs, fmt.Errorf("%v", errMsg))
	}
}

// Error returns a combined error message with all errors, or nil if no errors
func (h *ErrorHandler) Error() error {
	if len(h.errs) == 0 {
		return nil
	}

	var errorMsgs []string
	for _, err := range h.errs {
		if err != nil {
			errorMsgs = append(errorMsgs, err.Error())
		}
	}

	if len(errorMsgs) == 1 {
		return fmt.Errorf("%v", errorMsgs[0])
	}
	return fmt.Errorf("multiple errors:\n  - %s", strings.Join(errorMsgs, "\n  - "))
}

// HasErrors returns true if any errors were pushed
func (h *ErrorHandler) HasErrors() bool {
	return len(h.errs) > 0
}

// WrapError wraps an existing error with additional context
func (h *ErrorHandler) WrapError(prefix string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", prefix, err)
}
