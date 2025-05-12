#!/bin/bash
go run main.go --transport streamable-http \
--cert ./mcp.sea-flea.local.crt \
--key ./mcp.sea-flea.local.key \
--debug \
--demo-tools \
--demo-resources \
--demo-prompts

