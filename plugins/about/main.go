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

	resources := []Resource{
		{
			URI:         "about:///sea-flea",
			Name:        "resource_sample",
			Description: "This is a resource example",
			MimeType:    "application/json",
			//Text:        `{"message": "` + message + `", "version": "` + version + `", "project": "` + project + `", "id": "` + id + `"}`,
		},
	}
	jsonData, _ := json.Marshal(resources)
	pdk.OutputString(string(jsonData))

}

//go:export resource_sample
func ResourceSample() {
	message, okm := pdk.GetConfig("WASM_MESSAGE")
	if !okm {
		message = "..."
	}

	version, okv := pdk.GetConfig("WASM_VERSION")
	if !okv {
		version = "..."
	}

	project, okp := pdk.GetConfig("project")
	if !okp {
		project = "..."
	}
	id, okid := pdk.GetConfig("id")
	if !okid {
		id = "..."
	}

	resource := Resource{
		URI:         "about:///sea-flea",
		Name:        "resource_sample",
		Description: "This is a resource example",
		MimeType:    "application/json",
		Text:        `{"message": "` + message + `", "version": "` + version + `", "project": "` + project + `", "id": "` + id + `"}`,
	}

	jsonData, _ := json.Marshal(resource)
	pdk.OutputString(string(jsonData))

}

func main() {}
