package mcp

import (
	"sea-flea/config"
	"sea-flea/jsonrpc"
	"sea-flea/utils"
)

// func (s *MCPServer) handleResourcesList(params resources.ResourceListParams) (resources.ResourceListResult, *jsonrpc.JSONRPCError) {

// Handler for resources/list
func (s *MCPServer) handleResourcesList() (map[string]any, *jsonrpc.JSONRPCError) {
	if !s.initialized {
		return nil, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InvalidRequest,
			Message: "Server not initialized",
		}
	}

	resourcesList := make([]map[string]any, 0, len(s.resourceSet))

	for _, resource := range s.resourceSet {

		resourceMap := map[string]any{
			"name":        resource.Name,
			"description": resource.Description,
			"uri":         resource.URI,
			"mimeType":    resource.MimeType,
		}

		resourcesList = append(resourcesList, resourceMap)
	}

	output := map[string]any{
		"resources": resourcesList,
	}

	// Log the output
	utils.Log(func() string {
		jsonString, _ := utils.GenerateJsonStringFromMap(output)
		return "📝 resources/list\n" + jsonString
	}, config.LogOutput)

	return output, nil

}
