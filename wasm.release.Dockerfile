# 🚧 Work in progress
FROM  k33g/sea-flea:demo
WORKDIR /app

# Change this part
COPY plugins/*.wasm ./plugins/

ENTRYPOINT ["./sea-flea"]
