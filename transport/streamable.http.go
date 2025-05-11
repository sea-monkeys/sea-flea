package transport

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sea-flea/config"
	"sea-flea/jsonrpc"
	"sea-flea/mcp"
)

func StreamableHTTP(server *mcp.MCPServer) {

	var httpPort = os.Getenv("HTTP_PORT")
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

		log.Println("📝",
			"Received request: Method=", request.Method,
			"URL=", request.URL.String(),
			"Headers=", request.Header,
		)

		if request.Method != http.MethodPost {
			errorMessage := "Invalid request method"
			log.Println("😡", errorMessage)
			http.Error(response, errorMessage, http.StatusMethodNotAllowed)
			return
		}

		var jsonRPCRequest jsonrpc.JSONRPCRequest

		if err := json.NewDecoder(request.Body).Decode(&jsonRPCRequest); err != nil {
			errorMessage := "Invalid JSON-RPC request"
			log.Println("😡", errorMessage+":", err)
			http.Error(response, "😡 Invalid JSON-RPC request", http.StatusBadRequest)
			return
		}

		if jsonRPCRequest.JSONRPC != config.JSONRPCVersion {
			errorMessage := "Invalid JSON-RPC version"
			log.Println("😡", errorMessage)
			http.Error(response, errorMessage, http.StatusBadRequest)
			return
		}

		log.Printf("Received JSON-RPC request: %+v", jsonRPCRequest)

		jsonRPCResponse := server.HandleRequest(jsonRPCRequest)

		log.Printf("Generated JSON-RPC response: %+v", jsonRPCResponse)

		if jsonRPCResponse.Error != nil {
			errorMessage := "Error in JSON-RPC response"
			log.Println("😡", errorMessage+":", jsonRPCResponse.Error)
			http.Error(response, errorMessage, http.StatusBadRequest)
			return
		}

		if err := json.NewEncoder(response).Encode(jsonRPCResponse); err != nil {
			errorMessage := "Error in JSON-RPC response"
			log.Println("😡", errorMessage+":", err)
			http.Error(response, errorMessage, http.StatusInternalServerError)
		}

		//handlers.MainHandler(response, request)
	})

	var errListening error
	log.Println("🌍 Streamable HTTP MCP server is listening on: " + httpPort)
	errListening = http.ListenAndServe(":"+httpPort, mux)
	log.Fatalln("😡", errListening)
}
