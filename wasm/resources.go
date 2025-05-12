package wasm

import (
	"encoding/json"
	"fmt"
	"os"
	"sea-flea/mcp"
	"sea-flea/resources"

	extism "github.com/extism/go-sdk"
)

func registerResourcesOfThePlugin(server *mcp.MCPServer, pluginInst *extism.Plugin) {
	_, output, err := pluginInst.Call("resources_information", nil)
	if err != nil {
		// Handle error case
		fmt.Fprintf(os.Stderr, "Error calling resources_information: %v\n", err)
		return
	}

	var resourcesList []resources.Resource
	err = json.Unmarshal(output, &resourcesList)

	if err != nil {
		// Handle error case
		fmt.Fprintf(os.Stderr, "Error unmarshalling resources information: %v\n", err)
		return
	}

	// Register the resources with the server
	for _, resource := range resourcesList {
		// Define the resource handler
		// This function will be called when the resource is invoked
		// It will receive the arguments as a map[string]any
		// and should return the result as a string
		resource.ContentHandler = func(args map[string]any) (resources.ResourceContent, error) {
			return resources.ResourceContent{
				URI:      resource.URI,
				MimeType: resource.MimeType,
				Text:     resource.Text,
				Blob:     resource.Blob,
			}, nil
		}
		// Register the resource with the server
		server.AddResource(resource)
	}

}
