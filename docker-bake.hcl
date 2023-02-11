# vim: filetype=hcl softtabstop=2 tabstop=2 shiftwidth=2 fileencoding=utf-8 commentstring=#%s expandtab
# code: language=terraform insertSpaces=true tabSize=2
# ────────────────────────────────────────────────────────────────────────────────
#
# ╭──────────────────────────────────────────────────────────╮
# │                     variables                            │
# ╰──────────────────────────────────────────────────────────╯
# setting this variable to true will make stop the process from pushing the
# image to upstream docker registry.
#
# It is recommended to set this to true when working/experimenting with image
# builds; since by default images built with `bake` command do not show up when
# one runs `docker image ls`; `bake` command pushes the image to upstream
# registry and that can take a while depending on the image size; setting this
# variable to `true` makes sure that the built image is exported to local
# docker daemon so it would show up when one runs `docker image ls`.
variable "LOCAL" {default=false}
# hostname of the upstream registry that stores the main images.
# for this project, the main images are stored in dockerhub
variable "REGISTRY_HOSTNAME" {default="docker.io"}
# username in upstream registry that stores the main images.
variable "REGISTRY_USERNAME" {default="fjolsvin"}
# trigger arm64 builds
variable "ARM64" {default=true}
# trigger amd64 builds
variable "AMD64" {default=true}
# sets image tag. You can use the following
# environment variables to set this value:
#
# export TAG="$(git describe --tags --abbrev=0 2>/dev/null || true)"
variable "TAG" {default=""}
# ────────────────────────────────────────────────────────────
# default build group
#
# The targets in `default` group are built when no specific build target is
# passed to buildx; i.e
# ─── SNIPPETS ───────────────────────────────────────────────────────────────────
# docker buildx bake --builder "$(basename -s .git "$(git remote get-url origin)")"
group "default" {
  targets = [
    "release",
  ]
}
# ╭──────────────────────────────────────────────────────────╮
# │                   image build targets                    │
# ╰──────────────────────────────────────────────────────────╯
target "release" {
  context    = "."
  dockerfile = "contrib/docker/release/Dockerfile"
  platforms = [
    equal(AMD64,true) ? "linux/amd64":"",
    equal(ARM64,true) ? "linux/arm64":"",
  ]
  tags       = [
    equal(LOCAL,true)
    ? "podinfo"
    : equal("",TAG)
      ? ""
      : "${REGISTRY_HOSTNAME}/${REGISTRY_USERNAME}/podinfo:${TAG}",
  ]
  cache-from = [
    equal(LOCAL,true)
    ? ""
    : "type=registry,mode=max,ref=${REGISTRY_HOSTNAME}/${REGISTRY_USERNAME}/podinfo:cache" ,
  ]
  cache-to   = [
    equal(LOCAL,true)
    ? ""
    : "type=registry,mode=max,ref=${REGISTRY_HOSTNAME}/${REGISTRY_USERNAME}/podinfo:cache" ,
  ]
  output     = [
    equal(LOCAL,true)
    ? "type=docker"
    : "type=registry",
  ]
}
