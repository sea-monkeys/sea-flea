package mcp

import (
	"fmt"
	"log"
	"sea-flea/config"
	"sea-flea/jsonrpc"
	"sea-flea/prompts"
	"sea-flea/resources"
)

func (s *MCPServer) HandleRequest(request jsonrpc.JSONRPCRequest) jsonrpc.JSONRPCResponse {
	
	response := jsonrpc.JSONRPCResponse{
		JSONRPC: config.JSONRPCVersion,
		ID:      request.ID,
	}

	switch request.Method {
	case "initialize":
		result, err := s.handleInitialize()
		if err != nil {
			response.Error = err
		} else {
			response.Result = result
		}

	case "notifications/initialized":
		// Notifications don't require a response
		if err := s.handleInitialized(); err != nil {
			log.Printf("Error handling initialized notification: %v", err)
		}
		return jsonrpc.JSONRPCResponse{} // Empty response for notification

	case "tools/list":
		result, err := s.handleToolsList()
		if err != nil {
			response.Error = err
		} else {
			response.Result = result
		}

	case "tools/call":
		result, err := s.handleToolsCall(request.Params)
		if err != nil {
			response.Error = err
		} else {
			response.Result = result
		}

	case "ping":
		// Simple ping response
		response.Result = map[string]any{}

	case "completion/complete":
		log.Printf("Received completion notification")
		return jsonrpc.JSONRPCResponse{
			JSONRPC: config.JSONRPCVersion,
			ID:      request.ID,
			Result: map[string]interface{}{
				"completion": map[string]interface{}{
					"values":  []interface{}{},
					"hasMore": false,
				},
			},
		}

	//case "completion/cancel":
	//	log.Printf("Received completion/cancel notification")
	//	return jsonrpc.JSONRPCResponse{} // Empty response for notification

	case "resources/list":

		result, err := s.handleResourcesList()
		if err != nil {
			response.Error = err
		} else {
			response.Result = result
		}

	case "resources/read":
		var params resources.ResourceReadParams
		if request.Params == nil {
			//TODO
		}
		// get the URI value from request.Param
		params.URI = request.Params.(map[string]any)["uri"].(string)

		result, err := s.handleResourcesRead(params)
		if err != nil {
			response.Error = err
		} else {
			response.Result = result
		}

	case "prompts/list":
		result, err := s.handlePromptsList()
		if err != nil {
			response.Error = err
		} else {
			response.Result = result
		}

	case "prompts/get":
		var params prompts.PromptGetParams
		if request.Params == nil {
			//TODO
		}

		promptName, okName := request.Params.(map[string]any)["name"].(string)

		if !okName {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InvalidParams,
				Message: "Invalid parameters: name is required",
			}
			return response
		}

		promptArguments, okArguments := request.Params.(map[string]any)["arguments"].(map[string]any)
		if !okArguments {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InvalidParams,
				Message: "Invalid parameters: arguments is required",
			}
			return response
		}

		// Convert promptArguments to []map[string]any
		// This is a workaround for the MCP spec, which requires arguments to be an array of objects
		// where each object has a single key-value pair.
		// In this case, we convert the map to a slice of maps, each containing one key-value pair.
		// For example, {"arg1": "value1", "arg2": "value2"} becomes [{"arg1": "value1"}, {"arg2": "value2"}]
		convertedArguments := []map[string]any{}
		for key, value := range promptArguments {
			convertedArguments = append(convertedArguments, map[string]any{key: value})
		}

		params.Name = promptName
		params.Arguments = convertedArguments

		// Get the prompt from the prompt set
		result, err := s.handlePromptsGet(params)
		if err != nil {
			response.Error = err
		} else {
			response.Result = result
		}

	default:
		response.Error = &jsonrpc.JSONRPCError{
			Code:    jsonrpc.MethodNotFound,
			Message: fmt.Sprintf("Method not found: %s", request.Method),
		}
	}

	return response
}
