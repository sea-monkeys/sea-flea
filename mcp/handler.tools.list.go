package mcp

import (
	"sea-flea/jsonrpc"
	"sea-flea/utils"
)

//func (s *MCPServer) handleToolsList(params any) (any, *jsonrpc.JSONRPCError) {

//TODO: instead of any use structs

func (s *MCPServer) handleToolsList() (map[string]any, *jsonrpc.JSONRPCError) {
	if !s.initialized {
		return nil, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InvalidRequest,
			Message: "Server not initialized",
		}
	}

	toolsList := make([]map[string]any, 0, len(s.toolSet))
	// convert s.toolSet to a slice of maps

	// Each map contains the tool name, description, and input schema
	// The input schema is a map of argument names to their types
	// For example, {"arg1": "string", "arg2": "int"} becomes [{"arg1": "string"}, {"arg2": "int"}]
	// This is a workaround for the MCP spec, which requires arguments to be an array of objects
	// where each object has a single key-value pair.
	// Iterate over the tool set and create a list of tool information
	// Each tool information is a map with keys "name", "description", and "inputSchema"
	for _, tool := range s.toolSet {
		toolMap := map[string]any{
			"name":        tool.Name,
			"description": tool.Description,
			"inputSchema": tool.InputSchema,
		}
		toolsList = append(toolsList, toolMap)
	}

	output := map[string]any{
		"tools": toolsList,
	}

	// Log the output
	utils.Log(func() string {
		jsonString, _ := utils.GenerateJsonStringFromMap(output)
		return "📝 tools/list\n" + jsonString
	}, s.logOutput)

	return output, nil

}
