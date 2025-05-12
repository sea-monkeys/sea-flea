#!/bin/bash

extism call ../goodbye.wasm tools_information \
  --log-level "info" \
  --wasi
echo ""

extism call ../goodbye.wasm goodbye \
  --input '{"name":"Bob Morane"}' \
  --log-level "info" \
  --wasi
echo ""

extism call ../goodbye.wasm bye \
  --input '{"name":"Bob Morane"}' \
  --log-level "info" \
  --wasi
echo ""

extism call ../goodbye.wasm resources_information \
  --log-level "info" \
  --wasi
echo ""