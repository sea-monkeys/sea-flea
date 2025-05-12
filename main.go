package main

import (
	"os"
	"sea-flea/config"
	"sea-flea/demo"
	"sea-flea/mcp"
	"sea-flea/transport"
)

func main() {

	// Create server instance
	server := mcp.NewMCPServer()

	if config.LoadDemoTools {
		demo.LoadTools(server)
	}

	if config.LoadDemoResources {
		demo.LoadResources(server)
	}

	if config.LoadDemoPrompts {
		demo.LoadPrompts(server)
	}

	mcpMode := os.Getenv("MCP_MODE")
	if mcpMode == "" {
		mcpMode = "STDIO"
	}

	switch mcpMode {
	case "STREAMABLE_HTTP":
		transport.StreamableHTTP(server)
	case "STDIO":
		transport.STDIO(server)
	default:
		panic("Invalid MCP_MODE. Use 'STREAMABLE_HTTP' or 'STDIO'.")
	}

}
