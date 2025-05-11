package mcp

import (
	"fmt"
	"sea-flea/config"
	"sea-flea/jsonrpc"
	"sea-flea/prompts"
	"sea-flea/utils"
)

func (s *MCPServer) handlePromptsGet(params prompts.PromptGetParams) (map[string]any, *jsonrpc.JSONRPCError) {
	if !s.initialized {
		return nil, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InvalidRequest,
			Message: "Server not initialized",
		}
	}

	prompt, ok := s.GetPrompt(params.Name)
	if !ok {
		return map[string]any{}, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InvalidParams,
			Message: fmt.Sprintf("Prompt with name '%s' not found", params.Name),
		}
	}
	// convert params.Arguments to map[string]any
	// This is a workaround for the MCP spec, which requires arguments to be an array of objects
	// where each object has a single key-value pair.
	// In this case, we convert the map to a single map, where each key-value pair is merged.
	// For example, [{"arg1": "value1"}, {"arg2": "value2"}] becomes {"arg1": "value1", "arg2": "value2"}
	convertedArguments := map[string]any{}
	for _, arg := range params.Arguments {
		for key, value := range arg {
			convertedArguments[key] = value
		}
		// maps.Copy(convertedArguments, arg)
	}

	generatedPromptContents, err := prompt.ContentHandler(convertedArguments)
	if err != nil {
		return map[string]any{}, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InternalError,
			Message: fmt.Sprintf("Error generating content for prompt '%s': %v", params.Name, err),
		}
	}

	output := map[string]any{
		"messages": generatedPromptContents,
	}

	// Log the output
	utils.Log(func() string {
		jsonString, _ := utils.GenerateJsonStringFromMap(output)
		return "📝 prompts/get\n"+jsonString
	}, config.LogOutput)

	return output, nil

}
