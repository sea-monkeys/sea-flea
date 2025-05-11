package mcp

import (
	"sea-flea/config"
	"sea-flea/jsonrpc"
	"sea-flea/utils"
)

func (s *MCPServer) handlePromptsList() (any, *jsonrpc.JSONRPCError) {
	if !s.initialized {
		return nil, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InvalidRequest,
			Message: "Server not initialized",
		}
	}

	promptsList := make([]map[string]any, 0, len(s.promptSet))
	// Iterate over the prompt set and create a list of prompt information
	// Each prompt information is a map with keys "name", "description", and "arguments"
	// The "arguments" key is a list of maps, each containing a single key-value pair
	// representing the argument name and its type.
	// For example, {"arg1": "string", "arg2": "int"} becomes [{"arg1": "string"}, {"arg2": "int"}]
	// This is a workaround for the MCP spec, which requires arguments to be an array of objects
	// where each object has a single key-value pair.
	for name, prompt := range s.promptSet {
		promptInfo := map[string]any{
			"name":        name,
			"description": prompt.Description,
			"arguments":   prompt.Arguments,
		}
		promptsList = append(promptsList, promptInfo)
	}
	// Return in the correct format according to MCP spec
	output := map[string]any{
		"prompts": promptsList,
	}

	// Log the output
	utils.Log(func() string {
		jsonString, _ := utils.GenerateJsonStringFromMap(output)
		return "📝 prompts/list\n" + jsonString
	}, config.LogOutput)

	return output, nil
}
