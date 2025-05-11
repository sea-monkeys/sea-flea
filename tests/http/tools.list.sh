#!/bin/bash
: <<'COMMENT'
# Tools list
COMMENT

HTTP_PORT=5050
MCP_SERVER=http://localhost:${HTTP_PORT}
AUTHENTICATION_TOKEN=mcp-is-the-way

read -r -d '' DATA <<- EOM
{
  "jsonrpc": "2.0",
  "id": "1",
  "method": "tools/list",
  "params": {}
}
EOM

curl ${MCP_SERVER}/mcp \
  -H "Authorization: Bearer ${AUTHENTICATION_TOKEN}" \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -d "${DATA}" | jq
  #\
  #| grep "^data:" | sed 's/^data: //' | jq '.'

