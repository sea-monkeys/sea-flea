package demo

import (
	"fmt"
	"sea-flea/mcp"
	"sea-flea/prompts"
)

func LoadPrompts(server *mcp.MCPServer) {

	// ------------------------------------------------
	// Define the prompts
	// ------------------------------------------------
	basicPrompt := prompts.Prompt{
		Name:        "basic_prompt",
		Description: "A basic prompt example",
		Arguments: []map[string]any{
			{
				"name":        "message",
				"description": "a message",
				"required":    true,
			},
		},
		ContentHandler: func(args map[string]any) ([]map[string]any, error) {
			message, _ := args["message"].(string)
			result := fmt.Sprintf("You said: %s", message)
			return []map[string]any{
				{
					"role": "user",
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
				"name":        "firstName",
				"description": "First name of the person",
				"required":    true,
			},
			{
				"name":        "lastName",
				"description": "Last name of the person",
				"required":    true,
			},
		},
		ContentHandler: func(args map[string]any) ([]map[string]any, error) {
			firstName, _ := args["firstName"].(string)
			lastName, _ := args["lastName"].(string)
			result := fmt.Sprintf("Hello %s %s", firstName, lastName)

			return []map[string]any{
				{
					"role": "user",
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
}
