#!/bin/bash

extism call ../hey.wasm tools_information \
  --log-level "info" \
  --wasi
echo ""

extism call ../hey.wasm hey \
  --input '{"name":"Bob Morane"}' \
  --log-level "info" \
  --wasi
echo ""

extism call ../hey.wasm prompts_information \
  --log-level "info" \
  --wasi
echo ""

extism call ../hey.wasm hey_prompt \
  --input '{"name":"Bob Morane"}' \
  --log-level "info" \
  --wasi
echo ""

extism call ../hey.wasm hi_prompt \
  --input '{"firstName":"Bob","lastName":"Morane"}' \
  --log-level "info" \
  --wasi
echo ""