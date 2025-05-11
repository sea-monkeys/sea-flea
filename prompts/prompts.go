package prompts

// Prompt represents a single prompt definition
type Prompt struct {
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	Arguments   []map[string]any `json:"arguments"`
	ContentHandler func(map[string]any) ([]map[string]any, error)
	//Schema      map[string]any   `json:"schema,omitempty"`
}

// PromptGetParams represents the parameters for getting a prompt
type PromptGetParams struct {
	Name      string           `json:"name"`
	Arguments []map[string]any `json:"arguments"`
}

