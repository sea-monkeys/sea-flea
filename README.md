# Sea Flea MCP Server
> Available transports:
> - STDIO
> - Streamable HTTP


## STDIO

### 🐳 Build it

```bash
docker build -t mcp-sea-flea:demo .
```

## Use it with a MCP client (like Claude.AI Desktop)

```bash
{
  "mcpServers": {
    "MCP_SEA_FLEA" :{
      "command": "docker",
      "args": [
        "run",
        "--rm",
        "-i",
        "mcp-sea-flea:demo",
        "--debug",
        "--demo-tools",
        "--demo-resources",
        "--demo-prompts",
        "--plugins",
        "./plugins"
      ]
    }
  }
}
```
> With inspector: `docker run --rm -i mcp-sea-flea:demo --debug --demo-tools --demo-resources --demo-prompts --plugins ./plugins`


## Streamable HTTP
> 🚧 work in progress

### Start

- activate the transport with the `--transport streamable-http` option
- default HTTP port: `5050`, use the `--http-port <PORT>` option to change it


### Bearer Token

```bash
MCP_TOKEN="mcp-is-the-way" sea-fleat --transport streamable-http \
--debug \
--demo-tools \
--demo-resources \
--demo-prompts
```

## CI and Tests
> 🚧 work in progress

```bash
docker compose --file ci.compose.yml up --build
```

### STDIO Tests examples

```bash
npx @modelcontextprotocol/inspector --cli go run main.go --method tools/list
npx @modelcontextprotocol/inspector --cli docker run --rm -i mcp-sea-flea:demo --method tools/list
```

```bash
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
cat /tmp/mcp_test_input.jsonl | go run main.go --transport stdio --debug --demo-tools --demo-resources --demo-prompts | jq -s '.'
rm /tmp/mcp_test_input.jsonl
```


### Streamable HTTP examples
> 🚧 work in progress


## In progress



```bash
docker run --rm -p 3001:3001 \
  -e HTTP_PORT=3001 \
  -e PLUGINS_PATH=./plugins \
  -e PLUGINS_DEFINITION_FILE=plugins.yml \
  -v "$(pwd)/plugins":/app/plugins \
  -e RESOURCES_PATH=./resources \
  -e RESOURCES_DEFINITION_FILE=resources.yml \
  -v "$(pwd)/resources":/app/resources \
  -e PROMPTS_PATH=./prompts \
  -e PROMPTS_DEFINITION_FILE=prompts.yml \
  -v "$(pwd)/prompts":/app/prompts \
  -e WASIMANCER_ADMIN_TOKEN=wasimancer-rocks \
  -e WASIMANCER_AUTHENTICATION_TOKEN=mcp-is-the-way \
  -e UPLOAD_PATH=./plugins/bucket \
  k33g/wasimancer:0.0.7 
```


```bash
{
  "mcpServers": {
    "MCP_SEA_FLEA" :{
      "command": "docker",
      "args": [
        "run",
        "--rm",
        "-i",
        "mcp-sea-flea:demo",
        "--debug",
        "--demo-tools",
        "--demo-resources",
        "--demo-prompts",
        "--plugins",
        "./plugins"
      ]
    }
  }
}
```

```
docker run --rm -i -p 5050:5050 -v ./plugins:/plugins k33g/sea-flea:demo --debug --transport streamable-http --plugins /plugins
```

### Run it with an external mapping
```
docker run --rm -i -p 5050:5050 -v ${LOCAL_WORKSPACE_FOLDER}/plugins:/plugins k33g/sea-flea:demo --debug --transport streamable-http --plugins /plugins
```

### Build your own distribution

<!-- TODO --> Explain why the image is secure
<!-- TODO --> Explain what to do if you are using devcontainer

go run main.go --transport streamable-http --debug --plugins ./plugins
