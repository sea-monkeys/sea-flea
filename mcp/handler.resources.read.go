package mcp

import (
	"sea-flea/config"
	"sea-flea/jsonrpc"
	"sea-flea/resources"
	"sea-flea/utils"
)

// Handler for resources/read
func (s *MCPServer) handleResourcesRead(params resources.ResourceReadParams) (map[string]any, *jsonrpc.JSONRPCError) {

	if !s.initialized {
		return map[string]any{}, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InvalidRequest,
			Message: "Server not initialized",
		}
	}

	// Get the resource from the resource set
	resource, ok := s.GetResource(params.URI)
	if !ok {
		return map[string]any{}, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InvalidParams,
			Message: "Resource not found",
		}
	}

	resourceContent, err := resource.ContentHandler(map[string]any{
		"uri": params.URI,
	})
	if err != nil {
		return map[string]any{}, &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InternalError,
			Message: "Error reading resource",
		}
	}

	output := map[string]any{
		//"contents": resourceContent,
		"contents": []resources.ResourceContent{resourceContent},
	}

	// Log the output
	utils.Log(func() string {
		jsonString, _ := utils.GenerateJsonStringFromMap(output)
		return "📝 resources/read\n" + jsonString
	}, config.LogOutput)

	return output, nil
}
