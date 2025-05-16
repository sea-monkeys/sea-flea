package mcp

import (
	"encoding/json"
	"sea-flea/cli"
	"sea-flea/prompts"
	"sea-flea/resources"
	"sea-flea/tools"
)

// MCP specific types
type InitializeParams struct {
	ProtocolVersion string       `json:"protocolVersion"`
	Capabilities    Capabilities `json:"capabilities"`
	ClientInfo      ClientInfo   `json:"clientInfo"`
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
	Tools     map[string]any `json:"tools"`
	Resources map[string]any `json:"resources"`
	Prompts   map[string]any `json:"prompts"`
}

// Server state
type MCPServer struct {
	initialized bool
	//tools       []Tool
	toolSet     map[string]tools.Tool
	resourceSet map[string]resources.Resource
	promptSet   map[string]prompts.Prompt

	logOutput bool

	pluginsPath string

	pluginsSettings string
}

func NewMCPServer(cfg *cli.Config) *MCPServer {
	return &MCPServer{
		toolSet:     make(map[string]tools.Tool),
		resourceSet: make(map[string]resources.Resource),
		promptSet:   make(map[string]prompts.Prompt),

		logOutput:       cfg.Debug,
		pluginsPath:     cfg.PluginsPath,
		pluginsSettings: cfg.Settings,
	}
}

func (s *MCPServer) LogOutput() bool {
	return s.logOutput
}

func (s *MCPServer) PluginsPath() string {
	return s.pluginsPath
}

func (s *MCPServer) PluginsSettings() (map[string]string, error) {
	if s.pluginsSettings == "" {
		return nil, nil
	}
	// convert the json string to a map
	var settings map[string]string
	err := json.Unmarshal([]byte(s.pluginsSettings), &settings)
	if err != nil {
		return nil, err
	}
	return settings, nil
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

func (s *MCPServer) AddPrompt(prompt prompts.Prompt) {
	s.promptSet[prompt.Name] = prompt
}
func (s *MCPServer) GetPrompt(name string) (prompts.Prompt, bool) {
	prompt, ok := s.promptSet[name]
	return prompt, ok
}
