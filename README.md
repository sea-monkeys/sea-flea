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
        "mcp-sea-flea:demo"
      ]
    }
  }
}
```
> in inspector: docker run --rm -i mcp-sea-flea:demo

## Streamable HTTP
> 🚧 work in progress

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
cat /tmp/mcp_test_input.jsonl | go run main.go | jq -s '.'
rm /tmp/mcp_test_input.jsonl
```


### Streamable HTTP examples
> 🚧 work in progress
