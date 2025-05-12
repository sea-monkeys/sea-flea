package mcp

import (
	"encoding/json"
	"fmt"
	"sea-flea/jsonrpc"
	"sea-flea/tools"
	"sea-flea/utils"
)

func (s *MCPServer) handleToolsCall(params any) (map[string]any, *jsonrpc.JSONRPCError) {
	if !s.initialized {
		return nil, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InvalidRequest,
			Message: "Server not initialized",
		}
	}

	// Parse parameters
	paramsBytes, err := json.Marshal(params)
	if err != nil {
		return nil, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InternalError,
			Message: "Failed to parse parameters",
		}
	}

	var toolCall tools.ToolCallParams
	if err := json.Unmarshal(paramsBytes, &toolCall); err != nil {
		return nil, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InvalidParams,
			Message: "Invalid parameters",
		}
	}

	tool, ok := s.GetTool(toolCall.Name)
	if !ok {
		return nil, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.MethodNotFound,
			Message: fmt.Sprintf("Tool '%s' not found", toolCall.Name),
		}
	}

	return s.executeTool(tool.Name, toolCall.Arguments)

}

func (s *MCPServer) executeTool(toolName string, args map[string]any) (map[string]any, *jsonrpc.JSONRPCError) {

	result, err := s.toolSet[toolName].Handler(args)
	if err != nil {
		return nil, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InternalError,
			Message: fmt.Sprintf("Error executing tool '%s': %v", toolName, err),
		}
	}

	output := map[string]any{
		"content": []map[string]any{
			{
				"type": "text",
				"text": result.(string),
			},
		},
	}

	// Log the output
	utils.Log(func() string {
		jsonString, _ := utils.GenerateJsonStringFromMap(output)
		return "📝 tools/call\n" + jsonString
	}, s.logOutput)

	return output, nil

}
