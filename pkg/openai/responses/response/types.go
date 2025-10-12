package response

// Used at Used at /response/output/type
type OutputTypeEnum string

const (
	OutputTypeMessage        OutputTypeEnum = "message"
	OutputTypeFileSearchCall OutputTypeEnum = "file_search_call"
	OutputTypeFunctionCall   OutputTypeEnum = "function_call"
	OutputTypeWebSearchCall  OutputTypeEnum = "web_search_call"
	OutputTypeComputerCall   OutputTypeEnum = "computer_call"
	OutputTypeReasoning      OutputTypeEnum = "reasoning"
)

// Used at /response/output/status
type OutputStatusEnum string

const (
	OutputStatusInProgress OutputStatusEnum = "in_progress"
	OutputStatusCompleted  OutputStatusEnum = "completed"
	OutputStatusFailed     OutputStatusEnum = "failed"
)

type OutputBaseType interface {
	OutputType() OutputTypeEnum
}
