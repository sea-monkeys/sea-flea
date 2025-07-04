package main

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	"github.com/extism/go-pdk"
)

// -------------------------------------------------
//  Tools
// -------------------------------------------------

type Tool struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	//InputSchema map[string]any `json:"inputSchema"`
	InputSchema InputSchema `json:"inputSchema"`
}

type InputSchema struct {
	Type       string         `json:"type"`
	Required   []string       `json:"required"`
	Properties map[string]any `json:"properties"`
}

//go:export tools_information
func ToolsInformation() {
	roolDices := Tool{
		Name:        "roll_dices",
		Description: "a tool to roll dices",
		InputSchema: InputSchema{
			Type:     "object",
			Required: []string{"numFaces", "numDices"},
			Properties: map[string]any{
				"numFaces": map[string]any{
					"type":        "number",
					"description": "number of faces on the dice",
				},
				"numDices": map[string]any{
					"type":        "number",
					"description": "number of dices to roll",
				},
			},
		},
	}

	orcGreetings := Tool{
		Name:        "orc_greetings",
		Description: "make greetings as an orc",
		InputSchema: InputSchema{
			Type:     "object",
			Required: []string{"name"},
			Properties: map[string]any{
				"name": map[string]any{
					"type":        "string",
					"description": "Name of the person to say bye to",
				},
			},
		},
	}

	tools := []Tool{roolDices, orcGreetings}

	jsonData, _ := json.Marshal(tools)
	pdk.OutputString(string(jsonData))
}

//go:export orc_greetings
func OrcGreetings() {
	type Arguments struct {
		Name string `json:"name"`
	}
	arguments := pdk.InputString()
	var args Arguments
	json.Unmarshal([]byte(arguments), &args)

	pdk.OutputString("Throm-ka " + args.Name)
}

func rollTheDices(numDices, numFaces int) []int {
	// Seed the random number generator with current time
	rand.Seed(time.Now().UnixNano())

	// Create a slice to store the results
	results := make([]int, numDices)

	// Roll each die
	for i := 0; i < numDices; i++ {
		// Generate a random number between 1 and numFaces
		results[i] = rand.Intn(numFaces) + 1
	}

	return results
}

//go:export roll_dices
func RollDices() int32 {
	type Arguments struct {
		NumFaces int `json:"numFaces"`
		NumDices int `json:"numDices"`
	}

	arguments := pdk.InputString()

	var args Arguments
	json.Unmarshal([]byte(arguments), &args)
	numFaces := args.NumFaces
	numDices := args.NumDices

	results := rollTheDices(numDices, numFaces)
	// Calculate the sum of all dice
	sum := 0
	for _, result := range results {
		sum += result
	}

	pdk.OutputString(strconv.Itoa(sum))
	return 0
}

// -------------------------------------------------
//	Resources
// -------------------------------------------------
type Resource struct {
	URI         string `json:"uri"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	MimeType    string `json:"mimeType,omitempty"`
	Text        string `json:"text,omitempty"`
	Blob        string `json:"blob,omitempty"`
}

//go:export resources_information
func ResourcesInformation() {
	// Define the resources information
	resources := []Resource{
		{
			URI:         "message:///about",
			Name:        "dnd_about",
			Description: "About object",
			//MimeType:    "application/json",
		},
		{
			URI:         "message:///help",
			Name:        "dnd_help",
			Description: "Help resource",
			//MimeType:    "text/plain",
			//Text:        `=== Help ===`,
		},
	}
	jsonData, _ := json.Marshal(resources)
	pdk.OutputString(string(jsonData))

}

//go:export dnd_about
func DndAbout() {
	// Define the about resource content
	about := Resource{
		URI:         "message:///about",
		Name:        "dnd_about",
		Description: "About object",
		MimeType:    "application/json",
		Text: `{
			"version": "1.0.0",
			"author": "@k33g",
			"license": "MIT",
			"text": "This is a simple about object"
		}`,
	}

	jsonData, _ := json.Marshal(about)
	pdk.OutputString(string(jsonData))
}

//go:export dnd_help
func DndHelp() {
	// Define the help resource content
	help := Resource{
		URI:         "message:///help",
		Name:        "dnd_help",
		Description: "Help resource",
		MimeType:    "text/plain",
		Text: `=== Help ===
This is a simple help resource for the DnD plugin.
It provides information about the available resources and tools.`,
	}

	jsonData, _ := json.Marshal(help)
	pdk.OutputString(string(jsonData))
}

// -------------------------------------------------
//	Prompts
// -------------------------------------------------
type Prompt struct {
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	Arguments   []map[string]any `json:"arguments"`
}

//go:export prompts_information
func PromptsInformation() {
	requestInformationPrompt := Prompt{
		Name:        "request_information_prompt",
		Description: "a prompt to request information",
		Arguments: []map[string]any{
			{
				"name":        "name",
				"description": "name of the person from whom information is requested",
				"type":        "string",
			},
		},
	}

	rollDicesPrompt := Prompt{
		Name:        "roll_dices_prompt",
		Description: "A roll dices prompt example",
		Arguments: []map[string]any{
			{
				"name":        "numFaces",
				"description": "number of faces on the dice",
				"type":        "string",
			},
			{
				"name":        "numDices",
				"description": "number of dices to roll",
				"type":        "string",
			},
		},
	}

	prompts := []Prompt{requestInformationPrompt, rollDicesPrompt}

	jsonData, _ := json.Marshal(prompts)
	pdk.OutputString(string(jsonData))
}

type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
type Message struct {
	Role    string  `json:"role"`
	Content Content `json:"content"`
}

//go:export request_information_prompt
func RequestInformationPrompt() {
	type Arguments struct {
		Name string `json:"name"`
	}
	arguments := pdk.InputString()
	var args Arguments
	json.Unmarshal([]byte(arguments), &args)

	messages := []Message{
		{
			Role: "user",
			Content: Content{
				Type: "text",
				Text: "Please provide information about " + args.Name,
			},
		},
	}

	jsonData, _ := json.Marshal(messages)
	pdk.OutputString(string(jsonData))
}

//go:export roll_dices_prompt
func RollDicesPrompt() {
	type Arguments struct {
		NumFaces string `json:"numFaces"`
		NumDices string `json:"numDices"`
	}

	arguments := pdk.InputString()
	var args Arguments
	json.Unmarshal([]byte(arguments), &args)

	messages := []Message{
		{
			Role: "user",
			Content: Content{
				Type: "text",
				Text: "Please roll " + args.NumDices + " dice(s) with " + args.NumFaces + " faces",
			},
		},
	}

	jsonData, _ := json.Marshal(messages)
	pdk.OutputString(string(jsonData))
}

func main() {}
