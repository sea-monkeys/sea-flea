#!/bin/bash

extism call ../dnd.wasm tools_information \
  --log-level "info" \
  --wasi
echo ""

extism call ../dnd.wasm orc_greetings \
  --input '{"name":"Bob Morane"}' \
  --log-level "info" \
  --wasi
echo ""

extism call ../dnd.wasm roll_dices \
  --input '{"numFaces":6, "numDices":3}' \
  --log-level "info" \
  --wasi
echo ""

extism call ../dnd.wasm resources_information \
  --log-level "info" \
  --wasi
echo ""