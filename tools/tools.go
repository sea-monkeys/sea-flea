package tools

// Tool types
type Tool struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	InputSchema map[string]any `json:"inputSchema"`
	Handler     func(map[string]any) (any, error)
}


/*
type ToolNoHandler struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	InputSchema map[string]any `json:"inputSchema"`
}
*/

type ToolListParams struct {
	Cursor string `json:"cursor,omitempty"`
}

/*
type ToolListResult struct {
	Tools      []ToolNoHandler  `json:"tools"`
	NextCursor *string `json:"nextCursor,omitempty"`
}
*/

type ToolCallParams struct {
	Name      string                 `json:"name"`
	Arguments map[string]any `json:"arguments"`
}

type ToolContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type ToolCallResult struct {
	Content []ToolContent `json:"content"`
	IsError bool          `json:"isError,omitempty"`
}