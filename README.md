# Northern-Labs Interview

## Table of contents

- [Overview][1]

- [Code Statistics][11]

- [Usage and Demo][10]

- [Build Systems][2]

  - [Build Systems : Overview][3]
  - [Build Systems : mage][4]
  - [Build Systems : just][5]
  - [Build Systems : docker buildx][6]

- [Roadmap][12]

## Overview

> TODO

## Code Statistics

```console
❯ tokei
===============================================================================
 Language            Files        Lines         Code     Comments       Blanks
===============================================================================
 Dockerfile              2          656          566           89            1
 Go                    243        17497        13468         2525         1504
 JavaScript              1           25           23            2            0
 JSON                    5          291          291            0            0
 Shell                   5          205          179           22            4
 Plain Text              2          147            0          147            0
 TOML                    1           57           57            0            0
 YAML                    9          496          480           13            3
-------------------------------------------------------------------------------
 Markdown                3         1670            0         1517          153
 |- BASH                 1            7            7            0            0
 (Total)                           1677            7         1517          153
===============================================================================
 Total                 271        21044        15064         4315         1665
===============================================================================
```

## Usage and Demo

- Build and start the server

```bash
just run
```

server output is stored in `./tmp/server/log` file. You can `tail` it to follow
logs

- test `GET /healthz` endpoint

```console
❯ just liveness-probe
{
  "status": "OK"
}
```

## Build Systems

### Build Systems : Overview

This repo uses [`just`][7] ,[`mage`][8] and Docker's [`buildx`][9] plugin :

- `mage` : used for running Go toolchain tasks; e.g building the binary
- `just` : used for running other common tasks; such as running built binary or
  bootstrapping the development environment
- docker `buildx` : used for building and storing build-cache of multi-arch
  docker images using instructions in `docker-bake.hcl` file.

### Build Systems : mage

The following mage targets are available:

```console
Targets:
  build*    cross-compile the binary for all supported platforms and if possible, compress the binary
  clean     remove built binaries
  deps      tidy go modules and and downloads the dependencies
  test      run all tests across all sub-directories once.

* default target
```

The most commonly used target is `build` as it is the main target for building
the binary. You can use `mage -d "build/go" -w . "build"` to run this target

### Build Systems : just

The following just recipes are available:

```console
Available recipes:
    bootstrap            # installs dependencies and prepares development environment
    b                    # alias for `bootstrap`
    bootstrap-bash       # install all bash toolings
    bootstrap-git        # installs necessary git tools and configures git
    bootstrap-go         # install all go toolings
    bootstrap-json
    bootstrap-markdown   # install all markdown toolings
    bootstrap-os-pkgs    # this target installs a collection of core os packages. supports (debian, arch, alpine)
    bootstrap-pre-commit # ensures tools for making sane commits are installed and initializes pre-commit
    pc                   # alias for `bootstrap-pre-commit`
    bootstrap-semver     # bootstrap semantic versioning toolings
    build-go             # cross-compile go binaries for all supported platforms
    build                # alias for `build-go`
    clean-go             # removes build binaries (bin/) and tmp/ directory in repo's root
    clean                # alias for `clean-go`
    commit               # help guide the developers make conventional commits. it is recommended to use this target to make new commits
    c                    # alias for `commit`
    default              # `default` target, i.e target execued when just is called without any arguments
    format               # run all formatters
    f                    # alias for `format`
    fmt                  # alias for `format`
    format-bash          # detect and format all bash scripts
    bash-fmt             # alias for `format-bash`
    shfmt                # alias for `format-bash`
    format-go            # format all go files
    go-fmt               # alias for `format-go`
    gofmt                # alias for `format-go`
    format-json          # detect and format all json files
    json-fmt             # alias for `format-json`
    format-just          # format and stage the justfile
    just-fmt             # alias for `format-just`
    generate-changelog   # generate markdown and pdf changelog files
    gc                   # alias for `generate-changelog`
    git-add              # uses fzf to list git changes and help developers stage them
    ga                   # alias for `git-add`
    git-fetch            # fetches latest changes from upstream and removes any local branches that have been deleted in upstream
    gf                   # alias for `git-fetch`
    kary-comments        # adds support for extra languages to Kary Comments VSCode extension
    kc                   # alias for `kary-comments`
    kill                 # send SIGTERM to running binary to stop it
    lint                 # run all linters
    lint-bash            # lint all shellscripts
    shellcheck           # alias for `lint-bash`
    lint-go              # run golangci-lint with repo specific config
    golangci-lint        # alias for `lint-go`
    major-release        # generate changelog and create and push a new major release tag
    mar                  # alias for `major-release`
    minor-release        # generate changelog and create and push a new minor release tag
    patch-release        # generate changelog and create and push a new patch release tag
    pr                   # alias for `patch-release`
    run                  # build and start the server and forward logs to ./tmp/server/log
    snapshot             # take a tarball 'snapshot' of the repository.
    vscode-tasks         # generate vscode tasks.json file from justfile
    vt                   # alias for `vscode-tasks`
```

You can see a list of available recipes by running `just --list` or `just -l`.

The following are the most commonly used recipes:

- `just bootstrap` : installs all dependencies and prepares the development
  environment. Run it once after cloning the repo.
- `just build` : an alias for `mage build`
- `just run` : builds and runs the binary and forwards logs to
  `./tmp/server/log`
- `just lint` : runs all linters . Currently, recipes for `go` and `bash` are
  available
- `just fmt` : runs all formatters. Currently, recipes for `go`,`bash`,`json`
  and `justfile` are available

### Build Systems : docker buildx

- First and foremost , use must create a builder for this repo

```bash
docker buildx create --use --name "$(basename -s ".git" "$(git remote get-url origin)")" --driver docker-container
```

- Run all builds without pushing to Dockerhub. This is good for local testing

```bash
LOCAL=true docker buildx bake --builder "$(basename -s ".git" "$(git remote get-url origin)")"
```

- Run all builds and push to Docker Registry. This is good for running in CI/CD
  pipelines. Before running this snippet, set `REGISTRY_HOSTNAME` and
  `REGISTRY_USERNAME` environment variables to match your own setup.

```bash
export REGISTRY_HOSTNAME="docker.io" ;
export REGISTRY_USERNAME="fjolsvin" ;
docker buildx bake --builder "$(basename -s ".git" "$(git remote get-url origin)")"
```

- Besides `LOCAL` `REGISTRY_HOSTNAME` and `REGISTRY_USERNAME` variables, the
  following can be used to furthermore customize the build process:
  - `ARM64` : set it to `false` to disable building `aarch64` images
  - `AMD64` : set it to `false` to disable building `x86_64` images
  - `TAG` : you can set a custom tag for this build. as an example, the
    following snippet will tag the image with the latest git tag

```bash
export TAG="$(git describe --tags --abbrev=0 2>/dev/null || true)"
```

## Roadmap

- \[x] chores

  - \[x] setup pre-commit hooks
  - \[x] Linter setup
  - \[x] create a `gitpod` dockerfile for quickly spinning up dev Enviornments
    in Gitpod

- \[x] Update API framework's codebase to support `go1.20` from `go.14`

- \[x] Setup Github actions workflows

  - \[x] build and test on push
  - \[x] create a release archive on tags
  - \[x] build multi-platform docker image of the server

- \[x] `swagger` yaml file

- \[x] multi-arch docker builds with `docker buildx`

- \[x] Task runners, build and release systems

  - \[x] leverage `mage` build system to automate
    - \[x] cross-compilation of binaries for all supported platforms (linux,
      windows, darwin))
    - \[x] recursively detect all unit-tests and run them
  - \[x] `go-releaser` config for creating release archives for all supported
    platforms
  - \[x] `Just` targets to automate everything from bootstrapping a development
    environment to running end-to-end API tests
  - \[x] `vscode` **task** 'aliases\` to just targets

- Endpoints

  - \[x] **GET** `/healthz`
    - \[x] Implementation
    - \[x] E2E test automation
  - \[x] **GET** `/readyz`
    - \[x] Implementation
    - \[x] E2E test automation
    - \[x] `swagger` config
  - \[x] **GET** `/readyz/enable`
    - \[x] Implementation
    - \[x] E2E test automation
    - \[x] `swagger` config
  - \[x] **GET** `/readyz/disable`
    - \[x] Implementation
    - \[x] E2E test automation
    - \[x] `swagger` config
  - \[x] **GET** `/env`
    - \[x] Implementation
    - \[x] E2E test automation
    - \[x] `swagger` config
  - \[x] **GET** `/headers`
    - \[x] Implementation
    - \[x] E2E test automation
    - \[x] `swagger` config
  - \[x] **GET** `/delay/{seconds}`
    - \[x] Implementation
    - \[x] E2E test automation
    - \[x] `swagger` config
  - \[ ] Redis Group : There is a minor issue in the code that initializes the
    server which prevents server startup when redis config is not passed. This
    will be addressed immediately
    - \[x] pre-flight redis connection check
    - \[ ] **POST** `/cache/{key}`
      - \[x] Implementation
      - \[ ] validation
      - \[ ] E2E test automation
      - \[x] `swagger` config
    - \[ ] **PUT** `/cache/{key}`
      - \[x] Implementation
      - \[ ] validation
      - \[ ] E2E test automation
      - \[x] `swagger` config
    - \[ ] **GET** `/cache/{key}`
      - \[x] Implementation
      - \[ ] validation
      - \[ ] E2E test automation
      - \[x] `swagger` config
    - \[ ] **DELETE** `/cache/{key}`
      - \[x] Implementation
      - \[ ] validation
      - \[ ] E2E test automation
      - \[x] `swagger` config

- \[ ] docker-compose file with Redis

- \[ ] `k3d` config for Kubernetes native testing of the binary

[1]: #overview
[2]: #build-systems
[3]: #build-systems--overview
[4]: #build-systems--mage
[5]: #build-systems--just
[6]: #build-systems--docker-buildx
[7]: https://github.com/casey/just
[8]: https://magefile.org
[9]: https://docs.docker.com/build/bake/file-definition/
[10]: #usage-and-demo
[11]: #code-statistics
[12]: #roadmap
