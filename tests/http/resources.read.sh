#!/bin/bash
: <<'COMMENT'
# Resources list
COMMENT

HTTP_PORT=5050
MCP_SERVER=http://0.0.0.0:${HTTP_PORT}
AUTHENTICATION_TOKEN=mcp-is-the-way

# host.docker.internal

read -r -d '' DATA <<- EOM
{
  "jsonrpc": "2.0",
  "id": "1",
  "method": "resources/read",
  "params": {
    "uri": "message:///help"
  }
}
EOM

curl ${MCP_SERVER}/mcp \
  -H "Authorization: Bearer ${AUTHENTICATION_TOKEN}" \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -d "${DATA}" | jq


