# Extism Rust PDK Plugin

## Create Extism plugin

```bash
mkdir calc && cd calc
extism generate plugin # choose rust
rm -rf .git
```
> see: https://github.com/extism/cli?tab=readme-ov-file#generate-a-plugin


**Build**:
```bash
cd calc
cargo build --release 
cp target/wasm32-unknown-unknown/release/calc.wasm ../
```

**Run**:
```
extism call calc.wasm add \
  --input '{"a":30, "b":12}' \
  --log-level "info" \
  --wasi
```

```
extism call calc.wasm tools_information \
  --log-level "info" \
  --wasi
```

