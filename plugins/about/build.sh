#!/bin/bash
tinygo build -scheduler=none --no-debug \
  -o ../goodbye.wasm \
  -target wasi main.go

