variable "REPO" {
  default = "k33g"
}

variable "TAG" {
  default = "demo-wasm-files"
}

group "default" {
  targets = ["sea-flea"]
}

target "sea-flea" {
  context = "."
  dockerfile = "wasm.release.Dockerfile"
  args = {}
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
  tags = ["${REPO}/sea-flea:${TAG}"]
}

# docker buildx bake --push --file wasm.release.docker-bake.hcl
