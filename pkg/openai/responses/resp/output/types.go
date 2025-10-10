package output

// Used at Used at /response/output/type
type TypeEnum string

const (
	TypeMessage        TypeEnum = "message"
	TypeFileSearchCall TypeEnum = "file_search_call"
	TypeFunctionCall   TypeEnum = "function_call"
	TypeWebSearchCall  TypeEnum = "web_search_call"
	TypeComputerCall   TypeEnum = "computer_call"
)

// Used at /response/output/status
type StatusEnum string

const (
	StatusInProgress StatusEnum = "in_progress"
	StatusCompleted  StatusEnum = "completed"
	StatusFailed     StatusEnum = "failed"
)

type BaseType interface {
	outputType() TypeEnum
}
