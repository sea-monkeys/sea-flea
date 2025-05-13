#!/bin/bash
#go run main.go --transport stdio --debug --demo-tools --demo-resources --demo-prompts --plugins ./plugins
WASM_MESSAGE="Sea Flea" WASM_VERSION="0.0.0" go run main.go --transport streamable-http \
--debug \
--plugins ./plugins
