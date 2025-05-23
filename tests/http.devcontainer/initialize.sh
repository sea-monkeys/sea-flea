#!/bin/bash
: <<'COMMENT'
# Initialize

MCP (Model Context Protocol) requires a proper initialization handshake 
before any other operations can be performed. 
This is by design - it ensures that both client and server agree 
on protocol version and capabilities before exchanging data.
COMMENT

echo "🚀 Initializing MCP server..."

# Step 1: Initialize the server
read -r -d '' INIT_DATA <<- EOM
{
  "jsonrpc": "2.0",
  "id": 0,
  "method": "initialize",
  "params": {
    "protocolVersion": "2024-11-05",
    "capabilities": {},
    "clientInfo": {
      "name": "http-test-client",
      "version": "1.0.0"
    }
  }
}
EOM

echo "Sending initialization request..."
curl ${MCP_SERVER}/mcp \
  -H "Authorization: Bearer ${AUTHENTICATION_TOKEN}" \
  -H "Content-Type: application/json" \
  -d "${INIT_DATA}" | jq


# Step 2: Send the initialized notification (optional but recommended)
read -r -d '' INITIALIZED_DATA <<- EOM
{
  "jsonrpc": "2.0",
  "method": "notifications/initialized"
}
EOM

curl ${MCP_SERVER}/mcp \
  -H "Authorization: Bearer ${AUTHENTICATION_TOKEN}" \
  -H "Content-Type: application/json" \
  -d "${INITIALIZED_DATA}" > /dev/null | jq

