package demo

import (
	"sea-flea/mcp"
	"sea-flea/resources"
)

func LoadResources(server *mcp.MCPServer) {

	// ------------------------------------------------
	// Define the resources
	// ------------------------------------------------
	greetingResource := resources.Resource{
		URI:         "message:///greeting",
		Name:        "Greeting",
		Description: "A simple greeting resource",
		MimeType:    "text/plain",
		ContentHandler: func(params map[string]any) (resources.ResourceContent, error) {
			// Simulate fetching content
			content := "Hello, this is a greeting resource!"
			return resources.ResourceContent{
				URI:      "message:///greeting",
				MimeType: "text/plain",
				Text:     content,
			}, nil
		},
	}

	informationResource := resources.Resource{
		URI:         "message:///information",
		Name:        "Information",
		Description: "A simple information resource",
		MimeType:    "text/plain",
		ContentHandler: func(params map[string]any) (resources.ResourceContent, error) {
			// Simulate fetching content
			content := "Hello, this is an information resource!"
			return resources.ResourceContent{
				URI:      "message:///information",
				MimeType: "text/plain",
				Text:     content,
			}, nil
		},
	}
	// the URI is the identifier for the resource
	server.AddResource(greetingResource)
	server.AddResource(informationResource)
}
