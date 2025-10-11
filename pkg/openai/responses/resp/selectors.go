package resp

import "driver/pkg/openai/responses/resp/output"

type OutputSelector interface {
	SelectOutput(output.BaseType) bool
}

// OutputSelectorFunc is a function type that implements OutputSelector
type OutputSelectorFunc func(output.BaseType) bool

// SelectOutput implements the OutputSelector interface
func (f OutputSelectorFunc) SelectOutput(o output.BaseType) bool {
	return f(o)
}

// Predefined selectors
var (
	SelectMessage OutputSelectorFunc = func(o output.BaseType) bool {
		return o.OutputType() == output.TypeMessage
	}

	SelectFileSearchCall OutputSelectorFunc = func(o output.BaseType) bool {
		return o.OutputType() == output.TypeFileSearchCall
	}

	SelectWebSearchCall OutputSelectorFunc = func(o output.BaseType) bool {
		return o.OutputType() == output.TypeWebSearchCall
	}

	SelectFunctionCall OutputSelectorFunc = func(o output.BaseType) bool {
		return o.OutputType() == output.TypeFunctionCall
	}

	SelectComputerCall OutputSelectorFunc = func(o output.BaseType) bool {
		return o.OutputType() == output.TypeComputerCall
	}
)

// Helper functions to create custom selectors

// ByStatus creates a selector that filters by status (works for types that have Status field)
func ByStatus(status output.StatusEnum) OutputSelectorFunc {
	return func(o output.BaseType) bool {
		// Type assertion to check if output has Status field
		type hasStatus interface {
			GetStatus() output.StatusEnum
		}
		if s, ok := o.(hasStatus); ok {
			return s.GetStatus() == status
		}
		return false
	}
}

// And combines multiple selectors with AND logic
func And(selectors ...OutputSelector) OutputSelectorFunc {
	return func(o output.BaseType) bool {
		for _, sel := range selectors {
			if !sel.SelectOutput(o) {
				return false
			}
		}
		return true
	}
}

// Or combines multiple selectors with OR logic
func Or(selectors ...OutputSelector) OutputSelectorFunc {
	return func(o output.BaseType) bool {
		for _, sel := range selectors {
			if sel.SelectOutput(o) {
				return true
			}
		}
		return false
	}
}

// Not negates a selector
func Not(selector OutputSelector) OutputSelectorFunc {
	return func(o output.BaseType) bool {
		return !selector.SelectOutput(o)
	}
}
