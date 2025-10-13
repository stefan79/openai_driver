package openai

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

type Base64Bytes struct {
	Data     []byte
	MimeType string // e.g., "image/png", "application/pdf"
}

func NewBase64Bytes(data []byte, mimeType string) *Base64Bytes {
	b := &Base64Bytes{Data: data, MimeType: mimeType}
	return b
}

func (b Base64Bytes) MarshalJSON() ([]byte, error) {
	if b.Data == nil {
		return []byte("null"), nil
	}

	encoded := base64.StdEncoding.EncodeToString(b.Data)
	mimeType := b.MimeType

	// Format as data URI
	dataURI := "data:" + mimeType + ";base64," + encoded
	return json.Marshal(dataURI)
}

func (b *Base64Bytes) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	// Check if it's a data URI
	if strings.HasPrefix(s, "data:") {
		// Extract MIME type and data
		parts := strings.SplitN(s, ",", 2)
		if len(parts) == 2 {
			header := parts[0]
			if mimeParts := strings.SplitN(header, ";", 2); len(mimeParts) > 0 {
				b.MimeType = strings.TrimPrefix(mimeParts[0], "data:")
			}
			s = parts[1]
		}
	}

	// Remove base64 header if present
	s = strings.TrimPrefix(s, "base64,")

	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}

	b.Data = decoded
	return nil
}
