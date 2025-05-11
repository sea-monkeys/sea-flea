package transport

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sea-flea/config"
	"sea-flea/jsonrpc"
	"sea-flea/mcp"
	"sea-flea/utils"
)

func STDIO(server *mcp.MCPServer) {
	// Set up logging to stderr
	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("♒️ [STDIO] Starting MCP server...")

	// Create JSON encoder/decoder for stdio
	decoder := json.NewDecoder(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)

	// Main event loop
	for {
		var request jsonrpc.JSONRPCRequest
		if err := decoder.Decode(&request); err != nil {
			if err == io.EOF {
				log.Println("Connection closed")
				break
			}
			log.Printf("Error decoding request: %v", err)
			utils.SendError(encoder, nil, jsonrpc.ParseError, "Failed to parse JSON")
			continue
		}

		// Pretty print for debugging
		requestBytes, _ := json.MarshalIndent(request, "", "  ")
		log.Printf("Received request: %s", requestBytes)

		// Validate JSON-RPC version
		if request.JSONRPC != config.JSONRPCVersion {
			utils.SendError(encoder, request.ID, jsonrpc.InvalidRequest, "Only JSON-RPC 2.0 is supported")
			continue
		}

		// Handle the request
		response := server.HandleRequest(request)

		// Send response (if not a notification)
		if err := utils.SendResponse(encoder, response); err != nil {
			log.Printf("Error sending response: %v", err)
		}
	}
}
