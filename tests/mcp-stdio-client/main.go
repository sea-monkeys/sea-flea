package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
)

func main() {
	mcpClient, err := client.NewStdioMCPClient(
		"docker",
		[]string{}, // Empty ENV
		"run",
		"--rm",
		"-i",
		"k33g/sea-flea:stdio-0.0.0",
	)
	if err != nil {
		log.Fatalf("😡 Failed to create client: %v", err)
	}
	defer mcpClient.Close()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Initialize the client
	fmt.Println("🚀 Initializing mcp client...")
	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-curl client 🌍",
		Version: "1.0.0",
	}

	initResult, err := mcpClient.Initialize(ctx, initRequest)
	if err != nil {
		log.Fatalf("Failed to initialize: %v", err)
	}
	fmt.Printf(
		"🎉 Initialized with server: %s %s\n\n",
		initResult.ServerInfo.Name,
		initResult.ServerInfo.Version,
	)

	// List Tools
	fmt.Println("🛠️ Available tools...")
	toolsRequest := mcp.ListToolsRequest{}
	tools, err := mcpClient.ListTools(ctx, toolsRequest)
	if err != nil {
		log.Fatalf("😡 Failed to list tools: %v", err)
	}
	for _, tool := range tools.Tools {
		fmt.Printf("- %s: %s\n", tool.Name, tool.Description)
		fmt.Println("  Arguments:", tool.InputSchema.Properties)
	}
	fmt.Println()

	// Calling tool
	fmt.Println("📣 calling add")
	toolRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	toolRequest.Params.Name = "add"
	toolRequest.Params.Arguments = map[string]any{
		"a": 10,
		"b": 32,
	}
	result, err := mcpClient.CallTool(ctx, toolRequest)
	if err != nil {
		log.Fatalln("😡 Failed to call the tool:", err)
	}
	// display the text content of result
	fmt.Println("📝 content of result:")
	content, ok := result.Content[0].(mcp.TextContent)
	if !ok {
		log.Fatalln("😡 Failed to cast content to TextContent")
	}
	fmt.Println("🤖", content.Text)

}
