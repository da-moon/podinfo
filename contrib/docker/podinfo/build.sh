#!/usr/bin/env bash
#-*-mode:sh;indent-tabs-mode:nil;tab-width:2;coding:utf-8-*-
# vim: tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
set -xeuo pipefail
# env vars
export DOCKER_BUILDKIT="1"
export BUILDKIT_PROGRESS="plain"

WD="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../../" && pwd)"
ESC_WD="$(echo "$WD" | sed 's/\//\\\//g')"
DOCKER_FILE="$(readlink -f "$(dirname "${BASH_SOURCE[0]}")/Dockerfile")"
DOCKER_FILE=$(echo "${DOCKER_FILE}" | sed -e "s/$ESC_WD\///g")
pushd "$WD" >/dev/null 2>&1
if [ -z ${IMAGE_NAME+x} ] || [ -z ${IMAGE_NAME+x} ]; then
  IMAGE_NAME="docker.io/fjolsvin/$(basename "$(dirname "${BASH_SOURCE[0]}")")"
fi
CACHE_NAME="${IMAGE_NAME}:cache"
# Buildkit setup
[ -r "${HOME}/.docker/cli-plugins" ] && chmod a+x "${HOME}/.docker/cli-plugins/docker-buildx"
# Install Buildx if buildkit is enabled but the plugin does not exists
if [ -n "${DOCKER_BUILDKIT+x}" ] && [ -n "${DOCKER_BUILDKIT}" ]; then
  if [ "$DOCKER_BUILDKIT" = "1" ]; then
    if [[ ! $(docker buildx version 2>/dev/null) ]] &&
      ! sudo grep -sq 'docker\|lxc' '/proc/1/environ'; then
      if [ -n "${BUILDX_VERSION+x}" ] && [ -n "${BUILDX_VERSION}" ]; then
        DOWNLOAD_URL="https://github.com/docker/buildx/releases/download/v${BUILDX_VERSION}/buildx-v${BUILDX_VERSION}.linux-amd64"
      fi
      if [ -z ${DOWNLOAD_URL+x} ] || [ -z "${DOWNLOAD_URL}" ]; then
        REPO="docker/buildx"
        architecture="$(uname -m)"
        case "$architecture" in
          x86_64 | amd64)
            architecture="amd64"
            ;;
          aarch64)
            architecture="arm64"
            ;;
          *)
            echo >&2 "[ WARN ] unsopported architecture: $architecture"
            exit 0
            ;;
        esac
        DOWNLOAD_URL="$(curl -sL "https://api.github.com/repos/${REPO}/releases/latest" |
          jq -r "\
.assets[]|select(\
.browser_download_url \
| (\
contains(\"${architecture}\") \
and contains(\"linux\") \
and (contains(\"json\") | not))).browser_download_url")"
      fi
      mkdir -p "${HOME}/.docker/cli-plugins"
      curl --silent -L \
        --output "${HOME}/.docker/cli-plugins/docker-buildx" \
        "${DOWNLOAD_URL}"
      chmod a+x "${HOME}/.docker/cli-plugins/docker-buildx"
    fi
  fi
fi
BUILD=("docker")
if [[ $(docker buildx version 2>/dev/null) ]] &&
  ! sudo grep -sq 'docker\|lxc' '/proc/1/environ'; then
  builder="$(basename -s ".git" "$(git remote get-url origin)")"
  docker buildx create --use --name
  BUILD+=("buildx" "build")
  BUILD+=("--platform" "linux/amd64,linux/arm64")
  BUILD+=("--cache-to" "type=registry,mode=max,ref=${CACHE_NAME}")
  BUILD+=("--push")
  docker buildx use "${builder}" || docker buildx create --use --name "${builder}"
else
  BUILD+=("build")
  BUILD+=("--pull")
fi
BUILD+=("--file" "${DOCKER_FILE}")
BUILD+=("--cache-from" "type=registry,ref=${CACHE_NAME}")
BUILD+=("--build-arg" "PODINFO_API_ADDR=0.0.0.0:8080")
BUILD+=("--tag" "${IMAGE_NAME}:latest")
BUILD+=("--tag" "${IMAGE_NAME}:nightly")
TAG="$(git describe --tags --abbrev=0 2>/dev/null || true)"
if [ -n "${TAG+x}" ] && [ -n "${TAG}" ]; then
  BUILD+=("--tag" "${IMAGE_NAME}:${TAG#"v"}")
fi
BUILD+=("${WD}")
"${BUILD[@]}"
if [[ $(docker buildx version 2>/dev/null) ]] &&
  ! sudo grep -sq 'docker\|lxc' /proc/1/environ; then
  docker buildx use default
else
  PUSH="docker push"
  PUSH+=" ${IMAGE_NAME}:latest"
  $PUSH
fi
popd >/dev/null 2>&1
