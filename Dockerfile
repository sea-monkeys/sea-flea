FROM golang:1.24.2-alpine AS builder
WORKDIR /app

COPY . .

RUN <<EOF
go mod tidy 
go build
EOF

FROM scratch
WORKDIR /app
COPY --from=builder /app/sea-flea .
ENTRYPOINT ["./sea-flea"]
