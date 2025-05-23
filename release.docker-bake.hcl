variable "REPO" {
  default = "k33g"
}

variable "TAG" {
  #default = "0.0.0"
  default = "demo"
}

group "default" {
  targets = ["sea-flea"]
}

target "sea-flea" {
  context = "."
  dockerfile = "release.Dockerfile"
  args = {}
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
  tags = ["${REPO}/sea-flea:${TAG}"]
}

# docker buildx bake --push --file release.docker-bake.hcl
