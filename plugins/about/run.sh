#!/bin/bash

extism call ../about.wasm resources_information \
  --log-level "info" \
  --wasi
echo ""

extism call ../about.wasm resource_sample \
  --log-level "info" \
  --wasi
echo ""

