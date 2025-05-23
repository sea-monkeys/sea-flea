# 🚧 Work in progress
FROM --platform=$BUILDPLATFORM golang:1.24.3-alpine AS builder
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

COPY . .

RUN <<EOF
go mod tidy 
GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build
EOF

FROM scratch
WORKDIR /app
COPY --from=builder /app/sea-flea .

ENTRYPOINT ["./sea-flea"]
