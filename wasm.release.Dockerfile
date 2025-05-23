# 🚧 Work in progress
FROM  k33g/sea-flea:0.0.0
WORKDIR /app

# Change this part
COPY plugins/*.wasm ./plugins/

ENTRYPOINT ["./sea-flea"]
