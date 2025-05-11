package main

import (
	"fmt"
	"os"
	"sea-flea/mcp"
	"sea-flea/prompts"
	"sea-flea/resources"
	"sea-flea/tools"
	"sea-flea/transport"
)

func main() {

	// Define the tools
	calculatorTool := tools.Tool{
		Name:        "add",
		Description: "Perform addition of two numbers",
		InputSchema: map[string]any{
			"type":     "object",
			"required": []string{"a", "b"},
			"properties": map[string]any{
				"a": map[string]any{
					"type":        "number",
					"description": "First operand",
				},
				"b": map[string]any{
					"type":        "number",
					"description": "Second operand",
				},
			},
		},
		Handler: func(args map[string]any) (any, error) {
			a, _ := args["a"].(float64)
			b, _ := args["b"].(float64)
			result := a + b
			return fmt.Sprintf("%f", result), nil
		},
	}

	helloTool := tools.Tool{
		Name:        "hello",
		Description: "Say hello to someone",
		InputSchema: map[string]any{
			"type":     "object",
			"required": []string{"firstName", "lastName"},
			"properties": map[string]any{
				"firstName": map[string]any{
					"type":        "string",
					"description": "First argument",
				},
				"lastName": map[string]any{
					"type":        "string",
					"description": "Second argument",
				},
			},
		},
		Handler: func(args map[string]any) (any, error) {
			firstName, _ := args["firstName"].(string)
			lastName, _ := args["lastName"].(string)
			result := "Hello " + firstName + " " + lastName
			return result, nil
		},
	}

	vulcanSaluteTool := tools.Tool{
		Name:        "vulcan_salute",
		Description: "Perform the Vulcan salute",
		InputSchema: map[string]any{
			"type":     "object",
			"required": []string{"firstName", "lastName"},
			"properties": map[string]any{
				"firstName": map[string]any{
					"type":        "string",
					"description": "First argument",
				},
				"lastName": map[string]any{
					"type":        "string",
					"description": "Second argument",
				},
			},
		},
		Handler: func(args map[string]any) (any, error) {
			firstName, _ := args["firstName"].(string)
			lastName, _ := args["lastName"].(string)
			result := "🖖 Live long and prosper " + firstName + " " + lastName
			return result, nil
		},
	}

	// Create server instance
	server := mcp.NewMCPServer()
	server.AddTool(calculatorTool)
	server.AddTool(helloTool)
	server.AddTool(vulcanSaluteTool)

	// ------------------------------------------------
	// Add resources to the server
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



	basicPrompt := prompts.Prompt{
		Name:        "basic_prompt",
		Description: "A basic prompt example",
		Arguments: []map[string]any{
			{
				"name": "message",
				"description": "a message",
				"required": true,
			},
		},
		ContentHandler: func(args map[string]any) ([]map[string]any, error) {
			message, _ := args["message"].(string)
			result := fmt.Sprintf("You said: %s", message)
			return []map[string]any{
				{
					"role":    "user",
					"content": map[string]any{
						"type": "text",
						"text": result,
					},
				},

			}, nil
		},
	}

	helloPrompt := prompts.Prompt{
		Name:        "hello_prompt",
		Description: "A hello prompt example",
		Arguments: []map[string]any{
			{
				"name": "firstName",
				"description": "First name of the person",
				"required": true,
			},
			{
				"name": "lastName",
				"description": "Last name of the person",
				"required": true,
			},
		},
		ContentHandler: func(args map[string]any) ([]map[string]any, error) {
			firstName, _ := args["firstName"].(string)
			lastName, _ := args["lastName"].(string)
			result := fmt.Sprintf("Hello %s %s", firstName, lastName)
			
			return []map[string]any{
				{
					"role":    "user",
					"content": map[string]any{
						"type": "text",
						"text": result,
					},
				},

			}, nil
		},
	}




	server.AddPrompt(basicPrompt)
	server.AddPrompt(helloPrompt)

	mcpMode := os.Getenv("MCP_MODE")
	if mcpMode == "" {
		mcpMode = "STDIO"
	}

	switch mcpMode {
	case "STREAMABLE_HTTP":
		transport.StreamableHTTP(server)
	case "STDIO":
		transport.STDIO(server)
	default:
		panic("Invalid MCP_MODE. Use 'STREAMABLE_HTTP' or 'STDIO'.")
	}

}
