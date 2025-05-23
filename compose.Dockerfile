FROM golang:1.24.3-alpine AS builder
WORKDIR /app

COPY . .

RUN <<EOF
go mod tidy 
go build
EOF

FROM scratch
WORKDIR /app
COPY --from=builder /app/sea-flea .
# 🚧 Work in progress
COPY plugins/*.wasm ./plugins/
#ENV WASM_MESSAGE="Hello from the host!"
#ENV WASM_VERSION="9.9.9"
ENTRYPOINT ["./sea-flea"]
