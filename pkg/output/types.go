package output

type Outputter interface {
	VVVMessage(string)
	VVMessage(string)
	VMessage(string)
	Message(string)

	Output(o any)
}
