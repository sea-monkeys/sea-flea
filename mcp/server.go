package mcp

import (
	"sea-flea/resources"
	"sea-flea/tools"
)

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
	Resources map[string]any `json:"resources"`
}

// Server state
type MCPServer struct {
	initialized bool
	//tools       []Tool
	toolSet map[string]tools.Tool
	resourceSet map[string]resources.Resource
}

func NewMCPServer() *MCPServer {
	return &MCPServer{
		toolSet: make(map[string]tools.Tool),
		resourceSet: make(map[string]resources.Resource),
	}
}

func (s *MCPServer) AddTool(tool tools.Tool) {
	s.toolSet[tool.Name] = tool
}

func (s *MCPServer) GetTool(name string) (tools.Tool, bool) {
	tool, ok := s.toolSet[name]
	return tool, ok
}

func (s *MCPServer) AddResource(resource resources.Resource) {
	s.resourceSet[resource.URI] = resource
}

func (s *MCPServer) GetResource(uri string) (resources.Resource, bool) {
	resource, ok := s.resourceSet[uri]
	return resource, ok
}