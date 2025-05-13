package main

import (
	"encoding/json"

	"github.com/extism/go-pdk"
)

type Resource struct {
	URI         string `json:"uri"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	MimeType    string `json:"mimeType,omitempty"`
	Text        string `json:"text,omitempty"`
	Blob        string `json:"blob,omitempty"`
}

//export resources_information
func ResourcesInformation() {

	message, okm := pdk.GetConfig("WASM_MESSAGE")
	if !okm {
		message = "..."
	}

	version, okv := pdk.GetConfig("WASM_VERSION")
	if !okv {
		version = "..."
	}

	resources := []Resource{
		{
			URI:         "about:///sea-flea",
			Name:        "Resource sample",
			Description: "This is a resource example",
			MimeType:    "application/json",
			Text:        `{"message": "` + message + `", "version": "` + version + `"}`,
		},
	}
	jsonData, _ := json.Marshal(resources)
	pdk.OutputString(string(jsonData))

}

func main() {}
