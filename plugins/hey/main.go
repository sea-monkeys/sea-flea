package main

import (
	"encoding/json"

	"github.com/extism/go-pdk"
)

type Tool struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	InputSchema map[string]any `json:"inputSchema"`
}

//export tools_information
func tools_information() {
	tool := Tool{
		Name:        "hey",
		Description: "Say hey to someone",
		InputSchema: map[string]any{
			"type":     "object",
			"required": []string{"name"},
			"properties": map[string]any{
				"name": map[string]any{
					"type":        "string",
					"description": "Name of the person to say hey to",
				},
			},
		},
	}

	jsonData, _ := json.Marshal([]Tool{tool})
	pdk.OutputString(string(jsonData))
}

type Arguments struct {
	Name string `json:"name"`
}

//export hey
func hey() {
	arguments := pdk.InputString()
	var args Arguments
	json.Unmarshal([]byte(arguments), &args)
	pdk.OutputString("👋 hey " + args.Name + " 😆")
}

type Prompt struct {
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	Arguments   []map[string]any `json:"arguments"`
}

//export prompts_information
func prompts_information() {
	prompt1 := Prompt{
		Name:        "hey_prompt",
		Description: "A hey prompt example",
		Arguments: []map[string]any{
			{
				"name":        "name",
				"description": "Name of the person to say hey to",
				"required":    true,
			},
		},
	}

	prompt2 := Prompt{
		Name:        "hi_prompt",
		Description: "A hi prompt example",
		Arguments: []map[string]any{
			{
				"name":        "firstName",
				"description": "First name of the person",
				"required":    true,
			},
			{
				"name":        "lastName",
				"description": "Last name of the person",
				"required":    true,
			},
		},
	}

	jsonData, _ := json.Marshal([]Prompt{prompt1, prompt2})
	pdk.OutputString(string(jsonData))
}

//export hey_prompt
func hey_prompt() {
	arguments := pdk.InputString()
	type Arguments struct {
		Name string `json:"name"`
	}
	var args Arguments
	json.Unmarshal([]byte(arguments), &args)

	type Message struct {
		Role    string `json:"role"`
		Content struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	}
	
	messages := []Message{
		{
			Role: "user",
			Content: struct {
				Type string `json:"type"`
				Text string `json:"text"`
			}{
				Type: "text",
				Text: "📝👋 hey " + args.Name + " 😆",
			},
		},
		
	}

	jsonData, _ := json.Marshal(messages)
	pdk.OutputString(string(jsonData))

}

//export hi_prompt
func hi_prompt() {
	arguments := pdk.InputString()
	type Arguments struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}
	var args Arguments
	json.Unmarshal([]byte(arguments), &args)


	type Message struct {
		Role    string `json:"role"`
		Content struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	}
	
	results := []Message{
		{
			Role: "user",
			Content: struct {
				Type string `json:"type"`
				Text string `json:"text"`
			}{
				Type: "text",
				Text: "📝👋 hi " + args.FirstName + " " + args.LastName + " 😆",
			},
		},
		
	}

	jsonData, _ := json.Marshal(results)
	pdk.OutputString(string(jsonData))

}

func main() {}
