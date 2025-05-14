#!/bin/bash
tinygo build -scheduler=none --no-debug \
  -o ../about.wasm \
  -target wasi main.go

