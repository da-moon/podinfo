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
variable "AMD64" {default=true}
