# Sea Flea MCP Server

**Sea Flea** is a MCP Server "WASM Runner"

Release `0.0.0`, 🐳 Docker image: `k33g/sea-flea:0.0.0`

## Overview

Sea Flea is an MCP (Model Context Protocol) server that supports WebAssembly (WASM) plugins. Plugins can provide three types of capabilities:

- **Tools**: Functions that can be called with arguments
- **Resources**: Static or dynamic content accessible via URIs  
- **Prompts**: Templates for generating conversation prompts

And **Sea Flea** will load the plugin(s) and provide the JSON RPC endpoints.

> Available transports:
> - STDIO
> - Streamable HTTP
> WASM Runtime: [Extism](https://extism.org/) with the [Go SDK](https://github.com/extism/go-sdk) 

## Plugin Architecture

The Sea Flea server automatically discovers and loads plugins by:

1. Scanning for `.wasm` files in the plugins directory
2. Checking for required exported functions
3. Registering capabilities with the MCP server
4. Handling runtime calls to plugin functions

### Plugin Examples
<!-- TODO -->
👀 Have a look to the `./plugins` directory

### Key Concepts

#### Tools

Tools are interactive functions that can be called by the MCP client. Each tool must:

- **Define its schema** in `tools_information()` with name, description, and input schema
- **Have a corresponding exported function** with `//go:export function_name`
- **Accept JSON input** via `pdk.InputString()`
- **Return results** via `pdk.OutputString()`

**Tool Function Pattern:**
```go
//go:export your_tool_name
func YourToolName() {
    type Arguments struct {
        Param1 string `json:"param1"`
        Param2 int    `json:"param2"`
    }
    
    arguments := pdk.InputString()
    var args Arguments
    json.Unmarshal([]byte(arguments), &args)
    
    // Your logic here
    result := "Your result"
    
    pdk.OutputString(result)
}
```

#### Resources

Resources provide accessible content with URIs. They can be:

- **Static content** (documentation, templates, help text)
- **Dynamic content** based on configuration
- **Various formats** (JSON, text, markdown, binary data)

> The dynamic resources are not yet implemented with **Sea Flea**


**Resource URI Patterns:**
- `your-plugin:///resource-name`
- `config://setting-name`
- `message:///content-type`

#### Prompts

Prompts are templates that generate conversation messages. They:

- **Define required arguments** with descriptions and types
- **Generate structured message content** with roles and content
- **Return an array of messages** for conversation flow

**Prompt Function Pattern:**
```go
//go:export your_prompt_name
func YourPromptName() {
    type Arguments struct {
        Param1 string `json:"param1"`
    }
    
    arguments := pdk.InputString()
    var args Arguments
    json.Unmarshal([]byte(arguments), &args)

    messages := []Message{
        {
            Role: "user",
            Content: Content{
                Type: "text",
                Text: "Your generated prompt text with " + args.Param1,
            },
        },
    }

    jsonData, _ := json.Marshal(messages)
    pdk.OutputString(string(jsonData))
}
```


## Run Sea Flea Server (from code, with `go run`)

**STDIO Mode:**
```bash
go run main.go --transport stdio --debug --plugins ./plugins
```

**HTTP Mode:**
```bash
go run main.go --transport streamable-http --debug --plugins ./plugins
```

**With Environment Variables:**
```bash
WASM_MESSAGE="Custom message" \
WASM_VERSION="1.0.0" \
go run main.go --transport streamable-http --plugins ./plugins
```
> ✋ the environment variables names must start with `WASM_`

**With Plugin Settings:**
```bash
go run main.go --plugins ./plugins --settings '{"difficulty":"hard", "campaign":"storm"}'
```

### Configuration and Environment from the Plugins side

#### Environment Variables

Plugins can access environment variables starting with `WASM_`:

```go
message, ok := pdk.GetConfig("WASM_MESSAGE")
version, ok := pdk.GetConfig("WASM_VERSION")
```

#### Plugin Settings

Access settings passed via `--settings` flag:

```go
project, ok := pdk.GetConfig("project")
difficulty, ok := pdk.GetConfig("difficulty")
```

#### Plugin Filtering

Filter which plugins to load:

```bash
# not yet implemented - in progress
```

## Testing

### STDIO
<!-- TODO -->
👀 Have a look to the `./tests/stdio` directory

### HTTP
<!-- TODO -->
👀 Have a look to the `./tests/http` directory


## Docker Packaging

### Create the **sea-flea** base image:

```bash
docker buildx bake --push --file release.docker-bake.hcl
```


### Create an **sea-flea** image with you WASM plugins
> `wasm.release.Dockerfile`
```Dockerfile
FROM  k33g/sea-flea:0.0.0
WORKDIR /app

# Change this part
COPY plugins/*.wasm ./plugins/

ENTRYPOINT ["./sea-flea"]
```

Build it (with **Docker Bake** it's easier):
```hcl
variable "REPO" {
  default = "k33g"
}

variable "TAG" {
  default = "demo-wasm-files"
}

group "default" {
  targets = ["sea-flea"]
}

target "sea-flea" {
  context = "."
  dockerfile = "wasm.release.Dockerfile"
  args = {}
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
  tags = ["${REPO}/sea-flea:${TAG}"]
}
```

Then run:

```bash
docker buildx bake --push --file wasm.release.docker-bake.hcl
```

### Use it with Claude.AI Desktop (STDIO mode)

**Configuration**:
```json
{
    "MCP_SEA_FLEA_DEMO":{
        "command":"docker","args":[
          "run", 
          "--rm", 
          "-i", 
          "k33g/sea-flea:demo-wasm-files",
          "--debug",
          "--plugins",
          "./plugins"
        ]
    }
  }
}
```

### Test the Sea Flea container with the STDIO transport
<!-- TODO -->
👀 Have a look to the `./tests/stdio.container` directory

**Example**: get the list of the tools
```bash
# Create a temporary input file with the correct sequence
cat > /tmp/mcp_test_input.jsonl << 'EOF'
{
    "jsonrpc": "2.0", 
    "id": 0, 
    "method": "initialize", 
    "params": {"protocolVersion": "2024-11-05", "capabilities": {}, "clientInfo": {"name": "test", "version": "1.0.0"}}
}
{
    "jsonrpc": "2.0", 
    "method": "notifications/initialized"
}
{
    "jsonrpc": "2.0", 
    "id": 2, 
    "method": "tools/list", 
    "params": {}
}
EOF

# Run the server with the input file
echo "---------------------------------------------------------"
echo "Running MCP server with proper initialization sequence..."
echo "---------------------------------------------------------"

# Pipe the input to the server and process output with jq
cat /tmp/mcp_test_input.jsonl | docker run --rm -i k33g/sea-flea:demo-wasm-files --debug --plugins ./plugins | jq -c '.' | jq -s '.'

# Clean up
rm /tmp/mcp_test_input.jsonl
```

### Test the Sea Flea container with the HTTP transport

Start the server with the HTTP transport (with Bearer Token):
```bash
 docker run --rm -i \
 -e MCP_TOKEN="mcp-is-the-way" \
 -p 5050:5050 \
 k33g/sea-flea:demo-wasm-files \
 --debug \
 --transport streamable-http \
 --plugins ./plugins

```

<!-- TODO -->
👀 Have a look to the `./tests/http` directory

#### If you are running in devcontainer
> Tricky part 🥵
The easiest way to run tests on the MCP Server with the HTTP transport in a container and when working with devcontainer is **Docker Compose**.

👀 Have a look to the `./tests/http.devcontainer` directory


## Best Practices

### Plugin Development

1. **Use descriptive names** for tools, resources, and prompts
2. **Provide detailed descriptions** in schemas
3. **Handle errors gracefully** with meaningful messages
4. **Use consistent URI patterns** for resources
5. **Test thoroughly** with various inputs

### Performance

1. **Keep functions lightweight** - avoid heavy computation
2. **Use efficient JSON marshaling**
3. **Cache expensive operations** when possible
4. **Minimize memory allocations**

### Security

1. **Validate all inputs** before processing
2. **Sanitize outputs** to prevent injection
3. **Use safe defaults** for configuration
4. **Limit resource access** appropriately
