package mcp

import "sea-flea/jsonrpc"

//func (s *MCPServer) handleToolsList(params any) (any, *jsonrpc.JSONRPCError) {


func (s *MCPServer) handleToolsList() (any, *jsonrpc.JSONRPCError) {
	if !s.initialized {
		return nil, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InvalidRequest,
			Message: "Server not initialized",
		}
	}

	toolsList := make([]map[string]any, 0, len(s.toolSet))
	// convert s.toolSet to a slice of maps

	for _, tool := range s.toolSet {
		toolMap := map[string]any{
			"name":        tool.Name,
			"description": tool.Description,
			"inputSchema": tool.InputSchema,
		}
		toolsList = append(toolsList, toolMap)
	}

	return map[string]any{
		"tools": toolsList,
	}, nil
}
