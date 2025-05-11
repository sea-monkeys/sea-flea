package mcp

import (
	"sea-flea/jsonrpc"
	"sea-flea/resources"
)

// Handler for resources/read
func (s *MCPServer) handleResourcesRead(params resources.ResourceReadParams) (resources.ResourceReadResult, *jsonrpc.JSONRPCError) {

	if !s.initialized {
		return resources.ResourceReadResult{}, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InvalidRequest,
			Message: "Server not initialized",
		}
	}

	// Get the resource from the resource set
	resource, ok := s.GetResource(params.URI)
	if !ok {
		return resources.ResourceReadResult{}, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InvalidParams,
			Message: "Resource not found",
		}
	}

	resourceContent, err := resource.ContentHandler(map[string]any{
		"uri": params.URI,
	})
	if err != nil {
		return resources.ResourceReadResult{}, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InternalError,
			Message: "Error reading resource",
		}
	}

	return resources.ResourceReadResult{
		Contents: []resources.ResourceContent{
			resourceContent,
		},
	}, nil
}
