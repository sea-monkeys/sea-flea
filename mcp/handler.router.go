package mcp

import (
	"fmt"
	"log"
	"sea-flea/config"
	"sea-flea/jsonrpc"
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
		//result, err := s.handleInitialize(request.Params)
		if err != nil {
			response.Error = err
		} else {
			response.Result = result
		}

	case "notifications/initialized":
		// Notifications don't require a response
		//if err := s.handleInitialized(request.Params); err != nil {
		if err := s.handleInitialized(); err != nil {
			log.Printf("Error handling initialized notification: %v", err)
		}
		return jsonrpc.JSONRPCResponse{} // Empty response for notification

	case "tools/list":
		//result, err := s.handleToolsList(request.Params)
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

	default:
		response.Error = &jsonrpc.JSONRPCError{
			Code:    jsonrpc.MethodNotFound,
			Message: fmt.Sprintf("Method not found: %s", request.Method),
		}
	}

	return response
}
