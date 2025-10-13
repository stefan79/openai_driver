package request

type ToolTypeEnum string

const (
	ToolTypeCodeInterpreter    ToolTypeEnum = "code_interpreter"
	ToolTypeFileSearch         ToolTypeEnum = "file_search"
	ToolTypeWebSearch          ToolTypeEnum = "web_search"
	ToolTypeComputerUsePreview ToolTypeEnum = "computer_use_preview"
	ToolTypeMCP                ToolTypeEnum = "mcp"
	ToolTypeImageGeneration    ToolTypeEnum = "image_generation"
)

type ToolTypleWebSearchSearchContextSizeEnum string

const (
	ToolTypleWebSearchSearchContextSizeLow    ToolTypleWebSearchSearchContextSizeEnum = "low"
	ToolTypleWebSearchSearchContextSizeMedium ToolTypleWebSearchSearchContextSizeEnum = "medium"
	ToolTypleWebSearchSearchContextSizeHigh   ToolTypleWebSearchSearchContextSizeEnum = "high"
)

// Used in /request/tools/type
type BaseToolType interface {
	ToolType() ToolTypeEnum
}

// Used in /request/tools[@type=web_search]
type ToolTypeWebSearchDef struct {
	Type              ToolTypeEnum                             `json:"type"`
	SearchContextSize *ToolTypleWebSearchSearchContextSizeEnum `json:"search_context_size,omitempty"`
	Filters           *ToolTypeWebSearchFiltersDef             `json:"filters,omitempty"`
	UserLocation      *ToolTypeWebSearchSearchLocationDef      `json:"user_location,omitempty"`
}

func (t ToolTypeWebSearchDef) ToolType() ToolTypeEnum { return ToolTypeWebSearch }

type ToolTypeWebSearchFiltersDef struct {
	AllowedDomains []string `json:"allowed_domains,omitempty"`
}

type ToolTypeWebSearchSearchLocationDef struct {
	City     *string `json:"city,omitempty"`
	Country  *string `json:"country,omitempty"`
	Region   *string `json:"region,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	Type     *string `json:"type,omitempty"`
}
