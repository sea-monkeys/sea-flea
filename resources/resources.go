package resources

type ResourceListParams struct {
	Cursor *string `json:"cursor,omitempty"`
}

type Resource struct {
	URI            string `json:"uri"`
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
	MimeType       string `json:"mimeType,omitempty"`
	ContentHandler func(map[string]any) (ResourceContent, error)
}

type ResourceListResult struct {
	Resources []Resource `json:"resources"`
	IsError   bool       `json:"isError,omitempty"`

	// NextCursor *string    `json:"nextCursor,omitempty"`
}

type ResourceReadParams struct {
	URI string `json:"uri"`
}

type ResourceContent struct {
	URI      string `json:"uri"`
	MimeType string `json:"mimeType,omitempty"`
	Text     string `json:"text,omitempty"`
	Blob     string `json:"blob,omitempty"` // Base64 encoded for binary data
}

type ResourceReadResult struct {
	Contents []ResourceContent `json:"contents"`
}
