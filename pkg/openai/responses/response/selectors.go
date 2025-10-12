package response

type OutputSelector interface {
	SelectOutput(OutputBaseType) bool
}

// OutputSelectorFunc is a function type that implements OutputSelector
type OutputSelectorFunc func(OutputBaseType) bool

// SelectOutput implements the OutputSelector interface
func (f OutputSelectorFunc) SelectOutput(o OutputBaseType) bool {
	return f(o)
}

// Predefined selectors
var (
	SelectMessage OutputSelectorFunc = func(o OutputBaseType) bool {
		return o.OutputType() == OutputTypeMessage
	}

	SelectFileSearchCall OutputSelectorFunc = func(o OutputBaseType) bool {
		return o.OutputType() == OutputTypeFileSearchCall
	}

	SelectWebSearchCall OutputSelectorFunc = func(o OutputBaseType) bool {
		return o.OutputType() == OutputTypeWebSearchCall
	}

	SelectFunctionCall OutputSelectorFunc = func(o OutputBaseType) bool {
		return o.OutputType() == OutputTypeFunctionCall
	}

	SelectComputerCall OutputSelectorFunc = func(o OutputBaseType) bool {
		return o.OutputType() == OutputTypeComputerCall
	}
)

// Helper functions to create custom selectors

// ByStatus creates a selector that filters by status (works for types that have Status field)
func ByStatus(status OutputStatusEnum) OutputSelectorFunc {
	return func(o OutputBaseType) bool {
		// Type assertion to check if output has Status field
		type hasStatus interface {
			GetStatus() OutputStatusEnum
		}
		if s, ok := o.(hasStatus); ok {
			return s.GetStatus() == status
		}
		return false
	}
}

// And combines multiple selectors with AND logic
func And(selectors ...OutputSelector) OutputSelectorFunc {
	return func(o OutputBaseType) bool {
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
	return func(o OutputBaseType) bool {
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
	return func(o OutputBaseType) bool {
		return !selector.SelectOutput(o)
	}
}
