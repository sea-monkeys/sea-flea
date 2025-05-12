package transport

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sea-flea/config"
	"sea-flea/jsonrpc"
	"sea-flea/mcp"
	"sea-flea/utils"
	"strings"
)

func StreamableHTTP(server *mcp.MCPServer, cert, key string) {

	httpPort := os.Getenv("MCP_HTTP_PORT")
	if httpPort == "" {
		httpPort = "5050"
	}
	// Add token configuration
	authToken := os.Getenv("MCP_TOKEN")
	if authToken == "" {
		utils.Log(func() string {
			return "🔒 MCP_TOKEN environment variable should be set"
		}, server.LogOutput())
	}
	/*
		curl -X POST http://localhost:5050/mcp \
		-H "Authorization: Bearer your-secret-token-here" \
		-H "Content-Type: application/json" \
		-d '{"jsonrpc": "2.0", ...}'
	*/

	mux := http.NewServeMux()

	mux.HandleFunc("POST /mcp", func(response http.ResponseWriter, request *http.Request) {

		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		response.Header().Set("Content-Type", "application/json")

		if request.Method == http.MethodOptions {
			response.WriteHeader(http.StatusOK)
			return
		}

		if authToken != "" {
			// Check Authorization header
			authHeader := request.Header.Get("Authorization")

			if !strings.HasPrefix(authHeader, "Bearer ") || strings.TrimPrefix(authHeader, "Bearer ") != authToken {

				errorMessage := "Unauthorized: Invalid or missing Bearer token"

				utils.Log(func() string {
					return "🔒 " + errorMessage
				}, server.LogOutput())

				http.Error(response, errorMessage, http.StatusUnauthorized)
				return
			}
		}

		// Log the output
		utils.Log(func() string {
			return "🌍 Received HTTP Request:\n" +
				"  Method= " + request.Method + "\n" +
				"  URL= " + request.URL.String() + "\n" +
				"  Headers= " + request.Header.Get("Content-Type") + "\n" +
				"  Body= " + request.Header.Get("Content-Length") + "\n"
		}, server.LogOutput())

		if request.Method != http.MethodPost {
			errorMessage := "Invalid request method"

			utils.Log(func() string {
				return "😡 " + errorMessage
			}, server.LogOutput())

			http.Error(response, errorMessage, http.StatusMethodNotAllowed)
			return
		}

		var jsonRPCRequest jsonrpc.JSONRPCRequest

		if err := json.NewDecoder(request.Body).Decode(&jsonRPCRequest); err != nil {
			errorMessage := "Invalid JSON-RPC request"

			utils.Log(func() string {
				return "😡 " + errorMessage + ":" + err.Error()
			}, server.LogOutput())

			http.Error(response, errorMessage, http.StatusBadRequest)
			return
		}

		if jsonRPCRequest.JSONRPC != config.JSONRPCVersion {
			errorMessage := "Invalid JSON-RPC version"

			utils.Log(func() string {
				return "😡 " + errorMessage
			}, server.LogOutput())

			http.Error(response, errorMessage, http.StatusBadRequest)
			return
		}

		// Log the output
		utils.Log(func() string {
			// Pretty print for debugging
			requestBytes, _ := json.MarshalIndent(jsonRPCRequest, "", "  ")
			return "📶 Received JSON-RPC request:\n" + string(requestBytes)
		}, server.LogOutput())

		// 🧭 Handle the request and route it to the appropriate handler
		jsonRPCResponse := server.HandleRequest(jsonRPCRequest)

		utils.Log(func() string {
			// Pretty print for debugging
			requestBytes, _ := json.MarshalIndent(jsonRPCResponse, "", "  ")
			return "Ⓜ️ Generated JSON-RPC response:\n" + string(requestBytes)
		}, server.LogOutput())

		if jsonRPCResponse.Error != nil {
			errorMessage := "Error in JSON-RPC response"

			utils.Log(func() string {
				return "😡 " + errorMessage + ":" + jsonRPCResponse.Error.Message
			}, server.LogOutput())

			http.Error(response, errorMessage, http.StatusBadRequest)
			return
		}

		if err := json.NewEncoder(response).Encode(jsonRPCResponse); err != nil {
			errorMessage := "Error in JSON-RPC response"

			utils.Log(func() string {
				return "😡 " + errorMessage + ":" + err.Error()
			}, server.LogOutput())

			http.Error(response, errorMessage, http.StatusInternalServerError)
		}

		//handlers.MainHandler(response, request)
	})

	var errListening error

	if cert != "" && key != "" {
		log.Println("🔒 Streamable HTTP MCP server is listening on: " + httpPort)
		//log.Println("🔒 Using TLS with cert: " + cert + " and key: " + key)
		errListening = http.ListenAndServeTLS(":"+httpPort, cert, key, mux)
		log.Fatalln(errListening)
	} else {
		log.Println("🌍 Streamable HTTP MCP server is listening on: " + httpPort)
		errListening = http.ListenAndServe(":"+httpPort, mux)
		log.Fatalln(errListening)
	}

}
