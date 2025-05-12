#!/bin/bash
#go run main.go --transport stdio --debug --demo-tools --demo-resources --demo-prompts --plugins ./plugins
go run main.go --transport streamable-http \
--debug \
--plugins ./plugins

