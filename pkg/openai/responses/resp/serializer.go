package resp

import (
	"driver/pkg/openai/responses/resp/output"
)

type OutputSerializer interface {
	SelectOutput(*output.BaseType) []string
}

// OutputSerializerFunc is a function type that implements OutputSerializer
type OutputSerializerFunc func(*output.BaseType) []string

// SelectOutput implements the OutputSelector interface
func (f OutputSerializerFunc) SelectOutput(o *output.BaseType) []string {
	return f(o)
}

// Predefined selectors
var (
	SerializeText OutputSerializerFunc = func(o *output.BaseType) []string {
		if o == nil || (*o).OutputType() != output.TypeMessage {
			return nil
		}
		msg := (*o).(*output.TypeMessageDef)
		res := make([]string, 0)
		for _, content := range msg.Content {
			switch v := content.(type) {
			case *output.TypeMessageContentTypeOutputTextDef:
				res = append(res, v.Text)
			case *output.TypeMessageContentTypeRefusalDef:
				res = append(res, v.Refusal)
			}
		}
		return res
	}
)
