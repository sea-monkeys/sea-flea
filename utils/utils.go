package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"sea-flea/config"
	"sea-flea/jsonrpc"
)

func SendResponse(encoder *json.Encoder, response jsonrpc.JSONRPCResponse) error {
	if response.ID == nil {
		// Don't send empty responses for notifications
		return nil
	}

	// Pretty print for debugging
	responseBytes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return err
	}
	log.Printf("Sending response: %s", responseBytes)

	// Send the actual response
	return encoder.Encode(response)
}

func SendError(encoder *json.Encoder, id any, code int, message string) {
	response := jsonrpc.JSONRPCResponse{
		JSONRPC: config.JSONRPCVersion,
		ID:      id,
		Error: &jsonrpc.JSONRPCError{
			Code:    code,
			Message: message,
		},
	}
	if err := SendResponse(encoder, response); err != nil {
		log.Printf("Failed to send error response: %v", err)
	}
}

func GenerateJsonStringFromMap(output map[string]any) (string, error) {

	jsonString, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error generating JSON string: %v", err)
	}
	return string(jsonString), nil
}

func Log(callbck func() string, debug bool) {
	if debug {
		log.Println(callbck())
	}
}
