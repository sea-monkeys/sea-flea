package mcp

import "sea-flea/tools"

// MCP specific types
type InitializeParams struct {
	ProtocolVersion string      `json:"protocolVersion"`
	Capabilities    any `json:"capabilities"`
	ClientInfo      ClientInfo  `json:"clientInfo"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResult struct {
	ProtocolVersion string       `json:"protocolVersion"`
	ServerInfo      ServerInfo   `json:"serverInfo"`
	Capabilities    Capabilities `json:"capabilities"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Capabilities struct {
	Tools map[string]any `json:"tools"`
}

// Server state
type MCPServer struct {
	initialized bool
	//tools       []Tool
	toolSet map[string]tools.Tool
}

func NewMCPServer() *MCPServer {
	return &MCPServer{
		toolSet: make(map[string]tools.Tool),
	}
}

func (s *MCPServer) AddTool(tool tools.Tool) {
	s.toolSet[tool.Name] = tool
}
