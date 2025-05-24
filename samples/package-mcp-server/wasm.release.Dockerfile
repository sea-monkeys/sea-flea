# 🚧 Work in progress
# 👋 update the version/tag number
FROM  k33g/sea-flea:0.0.1
WORKDIR /app

# Adapt this part
COPY plugins/*.wasm ./plugins/

ENTRYPOINT ["./sea-flea"]

# Use this to run the container
# docker buildx bake --push --file wasm.release.docker-bake.hcl
