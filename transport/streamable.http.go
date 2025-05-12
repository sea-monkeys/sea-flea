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
)

func StreamableHTTP(server *mcp.MCPServer) {

	var httpPort = os.Getenv("MCP_HTTP_PORT")
	if httpPort == "" {
		httpPort = "5050"
	}
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

	log.Println("🌍 Streamable HTTP MCP server is listening on: " + httpPort)
	errListening := http.ListenAndServe(":"+httpPort, mux)
	log.Fatalln(errListening)
}
