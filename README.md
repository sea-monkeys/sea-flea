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
        "--demo-prompts"
      ]
    }
  }
}
```
> With inspector: `docker run --rm -i mcp-sea-flea:demo --debug --demo-tools --demo-resources --demo-prompts`




## Streamable HTTP
> 🚧 work in progress

- activate the transport with the `--transport streamable-http` option
- default HTTP port: `5050`, use the `--http-port <PORT>` option to change it

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
