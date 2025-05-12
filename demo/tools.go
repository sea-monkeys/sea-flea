package demo

import (
	"fmt"
	"sea-flea/mcp"
	"sea-flea/tools"
)

func LoadTools(server *mcp.MCPServer) {

	// ------------------------------------------------
	// Define the tools
	// ------------------------------------------------
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

	server.AddTool(calculatorTool)
	server.AddTool(helloTool)
	server.AddTool(vulcanSaluteTool)
}
