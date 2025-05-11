package mcp

import (
	"sea-flea/jsonrpc"
)

// func (s *MCPServer) handleResourcesList(params resources.ResourceListParams) (resources.ResourceListResult, *jsonrpc.JSONRPCError) {

// Handler for resources/list
func (s *MCPServer) handleResourcesList() (any, *jsonrpc.JSONRPCError) {
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

	return map[string]any{
		"resources": resourcesList,
	}, nil

}

/*
func (s *MCPServer) handleResourcesList() (resources.ResourceListResult, *jsonrpc.JSONRPCError) {

	resourcesList := make([]resources.Resource, 0, len(s.resourceSet))
	for _, resource := range s.resourceSet {

		resourceStruct := resources.Resource{
			Name:        resource.Name,
			Description: resource.Description,
			URI:         resource.URI,
			MimeType:    resource.MimeType,
		}
		resourcesList = append(resourcesList, resourceStruct)
	}


	return resources.ResourceListResult{
		Resources: resourcesList,
	}, nil

}

*/
