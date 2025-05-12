package wasm

import (
	"encoding/json"
	"fmt"
	"os"
	"sea-flea/mcp"
	"sea-flea/prompts"

	extism "github.com/extism/go-sdk"
)

func registerPromptsOfThePlugin(server *mcp.MCPServer, pluginInst *extism.Plugin) {
	_, output, err := pluginInst.Call("prompts_information", nil)
	if err != nil {
		// Handle error case
		fmt.Fprintf(os.Stderr, "Error calling prompts_information: %v\n", err)
		return
	}

	var promptsList []prompts.Prompt
	err = json.Unmarshal(output, &promptsList)
	if err != nil {
		// Handle error case
		fmt.Fprintf(os.Stderr, "Error unmarshalling prompts information: %v\n", err)
		return
	}
	// Register the prompts with the server
	for _, prompt := range promptsList {
		// Define the prompt handler
		// This function will be called when the prompt is invoked
		// It will receive the arguments as a map[string]any
		// and should return the result as a string
		prompt.ContentHandler = func(args map[string]any) ([]map[string]any, error) {
			// Call the plugin function with the prompt name and args
			// Convert args to JSON
			argsJSON, err := json.Marshal(args)
			if err != nil {
				// Handle error case
				fmt.Fprintf(os.Stderr, "Error marshalling args: %v\n", err)
				return nil, err
			}

			_, output, err := pluginInst.Call(prompt.Name, argsJSON)
			if err != nil {
				// Handle error case
				fmt.Fprintf(os.Stderr, "Error calling prompt %s: %v\n", prompt.Name, err)
				return nil, err
			}

			// Cast output to string
			result := string(output)

			fmt.Printf("🟩🟩🟩Result from plugin: %s\n", result)

			// Unmarshal the result to []map[string]any
			var resultList []map[string]any
			err = json.Unmarshal(output, &resultList)
			if err != nil {
				// Handle error case
				fmt.Fprintf(os.Stderr, "Error unmarshalling result: %v\n", err)
				return nil, err
			}
			// Return the result
			return resultList, nil


		}
		server.AddPrompt(prompt)
	}

}
