package response

type OutputSerializer interface {
	SelectOutput(*OutputBaseType) []string
}

// OutputSerializerFunc is a function type that implements OutputSerializer
type OutputSerializerFunc func(*OutputBaseType) []string

// SelectOutput implements the OutputSelector interface
func (f OutputSerializerFunc) SelectOutput(o OutputBaseType) []string {
	return f(&o)
}

// Predefined selectors
var (
	SerializeText OutputSerializerFunc = func(o *OutputBaseType) []string {
		if o == nil || (*o).OutputType() != OutputTypeMessage {
			return nil
		}
		msg := (*o).(*OutputTypeMessageDef)
		res := make([]string, 0)
		for _, content := range msg.Content {
			switch v := content.(type) {
			case *OutputTypeMessageContentTypeOutputTextDef:
				res = append(res, v.Text)
			case *OutputTypeMessageContentTypeRefusalDef:
				res = append(res, v.Refusal)
			}
		}
		return res
	}
)
