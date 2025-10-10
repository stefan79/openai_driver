package cli

import (
	"encoding/json"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

func DumpOutput(o any, f string, w io.Writer) error {
	switch f {
	case "json":
		return json.NewEncoder(w).Encode(o)
	case "yaml":
		return yaml.NewEncoder(w).Encode(o)
	default:
		return fmt.Errorf("unknown output format: %s", f)
	}
}
