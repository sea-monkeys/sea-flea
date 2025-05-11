package mcp

import (
	"encoding/json"
	"fmt"
	"sea-flea/jsonrpc"
	"sea-flea/tools"
)

func (s *MCPServer) handleToolsCall(params any) (any, *jsonrpc.JSONRPCError) {
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

func (s *MCPServer) executeTool(toolName string, args map[string]any) (any, *jsonrpc.JSONRPCError) {

	result, err := s.toolSet[toolName].Handler(args)
	if err != nil {
		return nil, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InternalError,
			Message: fmt.Sprintf("Error executing tool '%s': %v", toolName, err),
		}
	}
	// Check if the result is a string
	return tools.ToolCallResult{
		Content: []tools.ToolContent{
			{
				Type: "text",
				Text: result.(string),
			},
		},
	}, nil

}
