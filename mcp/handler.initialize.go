package mcp

import (
	"log"
	"sea-flea/config"
	"sea-flea/jsonrpc"
)

//func (s *MCPServer) handleInitialize(params any) (any, *jsonrpc.JSONRPCError) {

func (s *MCPServer) handleInitialize() (any, *jsonrpc.JSONRPCError) {
	s.initialized = true
	return InitializeResult{
		ProtocolVersion: config.ProtocolVersion,
		ServerInfo: ServerInfo{
			Name:    config.ServerName,
			Version: config.ServerVersion,
		},
		Capabilities: Capabilities{
			Tools: map[string]any{},
			Resources: map[string]any{
				"subscribe":   false, // optional
				"listChanged": false, // optional
			},
			Prompts: map[string]any{
				"supported": true,
				"list":      true,
				"get":       true,
			},
		},
	}, nil
}

//func (s *MCPServer) handleInitialized(params any) *jsonrpc.JSONRPCError {

func (s *MCPServer) handleInitialized() *jsonrpc.JSONRPCError {
	if !s.initialized {
		return &jsonrpc.JSONRPCError{
			Code:    jsonrpc.InvalidRequest,
			Message: "Server not initialized",
		}
	}
	log.Println("Server initialized successfully")
	return nil
}
