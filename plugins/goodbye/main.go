package main

import (
	"encoding/json"

	"github.com/extism/go-pdk"
)

type Tool struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	InputSchema map[string]any `json:"inputSchema"`
}

//export tools_information
func tools_information() {
	goodbyeTool := Tool{
		Name:        "goodbye",
		Description: "Say goodbye to someone",
		InputSchema: map[string]any{
			"type":     "object",
			"required": []string{"name"},
			"properties": map[string]any{
				"name": map[string]any{
					"type":        "string",
					"description": "Name of the person to say goodbye to",
				},
			},
		},
	}

	byeTool := Tool{
		Name:        "bye",
		Description: "Say bye to someone",
		InputSchema: map[string]any{
			"type":     "object",
			"required": []string{"name"},
			"properties": map[string]any{
				"name": map[string]any{
					"type":        "string",
					"description": "Name of the person to say bye to",
				},
			},
		},
	}

	tools := []Tool{goodbyeTool, byeTool}

	jsonData, _ := json.Marshal(tools)
	pdk.OutputString(string(jsonData))
}

type Resource struct {
	URI         string `json:"uri"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	MimeType    string `json:"mimeType,omitempty"`
	Text        string `json:"text,omitempty"`
	Blob        string `json:"blob,omitempty"`
}

//export resources_information
func resources_information() {
	// Define the resources information
	resources := []Resource{
		{
			URI:         "https://example.com/resource1",
			Name:        "Resource 1",
			Description: "This is the first resource",
			MimeType:    "application/json",
			Text: "This is the content of resource 1",
		},
		{
			URI:         "https://example.com/resource2",
			Name:        "Resource 2",
			Description: "This is the second resource",
			MimeType:    "text/plain",
			Text: "This is the content of resource 2",
		},
		{
			URI:         "https://example.com/resource3",
			Name:        "Resource 3",
			Description: "This is the third resource",
			MimeType:    "image/png",
			Text: "This is the content of resource 3",
			Blob: "This is the binary content of resource 3",
		},
	}
	jsonData, _ := json.Marshal(resources)
	pdk.OutputString(string(jsonData))

}

type Arguments struct {
	Name string `json:"name"`
}

//export goodbye
func goodbye() {
	arguments := pdk.InputString()
	var args Arguments
	json.Unmarshal([]byte(arguments), &args)
	pdk.OutputString("👋😢 Goodbye " + args.Name)

}

//export bye
func bye() {
	arguments := pdk.InputString()
	var args Arguments
	json.Unmarshal([]byte(arguments), &args)
	pdk.OutputString("👋👋👋 Bye " + args.Name)
}

func main() {}
