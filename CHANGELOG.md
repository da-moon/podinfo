# Changelog

## [v1.0.0][1] (2023-02-15)

### Documentation

- **changelog:** v0.9.0 8ead18d

### Build System

- **just/podinfo:** `/cache/` api group recipes 2108ed4

## [v0.9.0][2] (2023-02-15)

### Bug Fixes

- **api/handlers/cache:** all cache endpoints are working 63e63c7
- **api/handlers/cache:** enabled `/cache` api group 718eae8

### Documentation

- **changelog:** v0.9.0 9f50cc4

### Refactors

- better logging f61c835
- **build/go/version:** remove unnecessary information 6999c69

### Chore

- **docker-compose:** fix redis listener 3ad11d9
- **docker-compose:** `redis` service configuration file d3178be
- **docker-compose:** `redis` service 5f696a1

### [v0.8.2][3] (2023-02-15)

#### Documentation

- **changelog:** v0.9.0 868a12d

#### CI

- **docker:** checking to see if docker credenials exist 2e7d3fa
- **docker:** checking to see if docker credenials exist 06a0f86
- **docker:** trigger build after adding secrets e8be68d
- **docker:** fix `buildx` id collision 6f1f2a8

### [v0.8.1][4] (2023-02-15)

#### Documentation

- **changelog:** v0.8.0 9522755

#### CI

- **dockeer:** trigger build on changes to any `*.go` file 7ec1748

## [v0.8.0][5] (2023-02-15)

### Documentation

- add some instructions and context 344ff9a
- update `just` target in main `README` file 68e9779
- add `Roadmap` a3b0242
- `swagger` yaml efb738c
- **changelog:** v0.7.0 13a5032

## [v0.7.0][6] (2023-02-15)

### Features

- **cmd/podinfo/commands/server:** parse redis info with flags 8e57115
- **api/handlers/cache:** register **DELETE** `/cache/{key}` route in API group
  `Register()` a5ab3a8
- **api/handlers/cache:** register **DELETE** `/cache/{key}` route in API group
  `Register()` d12f7db
- **api/handlers/cache/delete:** `handler` implementation 07c94d5

### Documentation

- added code stats 3e0e1ef
- **api/handlers/cache/delete:** package synopsis d0e25cb
- **changelog:** v0.7.0 8b5e5d2

### [v0.6.3][7] (2023-02-15)

#### Features

- **api/handlers/cache:** register **GET** `/cache/{key}` route in API group
  `Register()` ab70334
- **api/handlers/cache/get:** route `register` implementation 9935643
- **api/handlers/cache/get:** `handler` implementation f154df8

#### Documentation

- **api/handlers/cache/doc:** package synopsis 9de9d35
- **changelog:** v0.7.0 61739b5

### [v0.6.2][8] (2023-02-15)

#### Features

- **api/handlers/cache:** register **PUT** `/cache/{key}` route in API group
  `Register()` b7cb2ba
- **api/handlers/cache/put:** route `register` implementation 6696e26
- **api/handlers/cache/put:** `handler` implementation 15a177c

#### Documentation

- **api/handlers/cache/put:** package synopsis 26257c2
- **changelog:** v0.7.0 1b6e7f1

### [v0.6.1][9] (2023-02-15)

#### Features

- **api/handlers/cache:** register **POST** `/cache/{key}` route in API group
  `Register()` f2f2164
- **api/handlers/cache/post:** route `register` implementation fb15f04
- **api/handlers/cache/post:** `handler` implementation 86f9e9d

#### Documentation

- **api/handlers/cache:** package synopsis 1e8dc7e
- **api/handlers/cache/post:** package synopsis f878c09
- **changelog:** v0.6.0 cf372a5

## [v0.6.0][10] (2023-02-15)

### Features

- **api/handlers/cache/shared:** `preflight` request handler 489db81
- **api/handlers/cache/shared:** `API group` shared constructs 47b687d
- **cmd/podinfo/commands/server:** flags to read `redis` config 76a07eb
- **api/core:** add `redis` configuration fields and methods cc54fcf
- **internal/cli/value:** handle `int` type cli flags d26c406

### Documentation

- **api/handlers/cache/shared:** package synopsis 2732306
- **changelog:** v0.5.0 c313777

### Refactors

- **internal/cli/flagset:** more flexibilty in prettifying generated flag usage
  893a1e2
- **internal/cli/flagset:** more flexiblity in self-documenting flags fc63a46

### Build System

- **docker/podinfo:** minor changes to build script 9e7f04a

### Chore

- **re-commit:** hook updates bff048a

## [v0.5.0][11] (2023-02-15)

### [v0.4.1][12] (2023-02-15)

#### Features

- **api/handlers:** register **GET** `/delay/{seconds}` route in `Initialize()`
  5a9dc8e
- **api/handlers/delay:** route `register` implementation f39830b
- **api/handlers/delay:** `handler` implementation bf469a3
- **api/handlers/delay:** `response` setup d22dc78

#### Documentation

- **changelog:** v0.5.0 c313777
- **api/handlers/delay:** package synopsis e260245
- **changelog:** v0.4.0 d7d5af7

#### Build System

- **build/just/podinfo:** `delay-probe` recipe 9d0b55f

## [v0.4.0][13] (2023-02-15)

### Features

- **api/handlers:** register `/headers` route in `Initialize()` 10cef5d
- **api/handlers/headers:** route `register` implementation 989f948
- **api/handlers/headers:** `handler` implementation 9871e64
- **api/handlers/headers:** `response` type 019355e

### Documentation

- **api/handlers/headers:** package synopsis 489218a

### Build System

- **build/just/podinfo:** `headers-probe` recipe 239a061

## [v0.3.0][14] (2023-02-15)

### Features

- **api/handlers:** register `/env` route in `Initialize()` 2d98b6a
- **api/handlers/env:** routte `register` implementation 958bc47
- **api/handlers/env:** `handler` implementation 0ef5a8f
- **api/handlers/env:** `response` type 877b7e4

### Documentation

- **api/handlers/env:** package synopsis edc489e

### Build System

- **build/just/podinfo:** `env-probe` recipe b66f28b

### [v0.2.2][15] (2023-02-15)

#### Features

- **api/handlers:** register `/readyz/enable` route in `Initialize()` dead392
- **api/handlers/readiness/enable:** route init b56a400
- **api/handlers/readiness/enable:** handler setup 9356de0

#### Documentation

- **api/handlers/readiness/enable:** package synopsis faeaf50

#### Build System

- **build/just/podinfo:** `readiness-probe-enable` recipe 65c6272

### [v0.2.1][16] (2023-02-15)

#### Features

- **api/handlers:** register the route in `Initialize()` ae99b84
- **api/handlers/readiness/disable:** `handler.Register()` implementation
  a3851e6
- **api/handlers/readiness/disable:** `HandlerFn` implementation 6297492
- **api/handlers/readiness/disable:** `handler.GetLogger()` synopsis 3e7a28f
- **api/handlers/readiness/disable:** `handler.GetLogger()` implementation
  fcdb30b
- **api/handlers/readiness/disable:** `handler.SetLogger()` synopsis ce5be1d
- **api/handlers/readiness/disable:** `handler.SetLogger()` implementation
  b39a633
- **api/handlers/readiness/disable:** `handler` struct implementation 32376d6

#### Documentation

- **api/handlers/readiness/disable:** `HandlerFn` synopsis ac1c0ee
- **api/handlers/readiness/disable:** `handler` struct synopsis 5e8ae1e
- **api/handlers/readiness/disable:** package synopsis 4753ee2
- **changelog:** v0.2.0 2ec37fc

#### Build System

- **build/just/podinfo:** `readiness-probe-disable` recipe f0efeb5

## [v0.2.0][17] (2023-02-15)

### Features

- **api/handlers/readiness:** `handler.Register()` implementation 229f6ef
- **api/handlers/readiness:** `init()` implementation 316c88d
- **api/handlers/readiness:** `HandlerFn` implementation fa4e4e0
- **api/handlers/readiness:** `handler.GetStatus()` implementation 6975ceb
- **api/handlers/readiness:** `handler.SetStatus()` implementation 0ae4c75
- **api/handlers/readiness:** `handler.GetLogger()` synopsis 1fb4903
- **api/handlers/readiness:** `handler.GetLogger()` implementation 2af2f25
- **api/handlers/readiness:** `handler.SetLogger()` implementation a34f22d
- **api/handlers/readiness:** `handler` struct implementation 4a73558
- **api/handlers/readiness:** `Response` struct implementation 8ea8855
- **api/handlers/readiness:** `Status.String()` implementation f81e205
- **api/handlers/readiness:** `Status` enum c45a277
- **sdk/api/response:** `LogEntry()` implementation c62180d

### Documentation

- **build/just/podinfo:** `readiness-probe` recipe a989ce1
- **api/handlers/readiness:** `HandlerFn` synopsis 5b226c3
- **api/handlers/readiness:** `handler.GetStatus()` synopsis 7f85163
- **api/handlers/readiness:** `handler.SetStatus()` synopsis db54578
- **api/handlers/readiness:** `handler.SetLogger()` synopsis 2ac7d67
- **api/handlers/readiness:** `handler` struct synopsis 14fcb57
- **api/handlers/readiness:** `Response` struct synopsis 7c7ecff
- **api/handlers/readiness:** `Status.String()` synopsis 777320c
- **api/handlers/readiness:** `Status` enum 9788dfb
- **api/handlers/readiness:** package synopsis 92c17f3
- **sdk/api/response:** `LogEntry()` synopsis c9c5f7d
- **changelog:** v0.2.0 c46195b

### Refactors

- **api/handlers/liveness:** new pattern of writing `HandlerFn` d270881
- **sdk/api/response:** `Write` function 66fd066
- **api/handlers/liveness:** change route impl pattern 6efbf25

### Build System

- **build/just/podinfo:** `readiness-probe` recipe 14a6fe0

### Chore

- **gitpod:** fix `gitpod` shell config in dockerfile ad05fc9

### [v0.1.1][18] (2023-02-13)

#### Documentation

- **changelog:** `v0.1.0` e1bdb33

#### Build System

- **just/git:** prune refs in `git-fetch` recipe 69d5e05

## [v0.1.0][19] (2023-02-13)

### Features

- **sdk/api/response:** `WriteJSON()` function implementation ac8c96c
- **sdk/api/response:** `WriteSuccessfulJSONRaw()` function implementation
  f264be1
- **api/handlers:** remove `Prefix` from pre-flight handler ae48e9d
- **api/handlers:** Register and initialize `kubernetes-liveness-probe` route
  1c416f5
- **api/handlers/liveness:** `Initialize()` function 73548a0
- **api/handlers/liveness:** `New()` function 016f07c
- **api/handlers/liveness:** `Handler` struct 133a5a8
- **api/handlers/liveness:** package constants a0b8f71
- **api/handlers/liveness:** `Response` struct 9bc1341

### Bug Fixes

- **sdk/api/response:** address linter complaints a9eed4e

### Documentation

- **docker/buildx:** `gitpod` target 11b8055
- `liveness-probe` instruction fc5852c
- **sdk/api/response:** `WriteJSON()` function synopsis 7e25e0a
- **sdk/api/response:** `WriteSuccessfulJSONRaw()` synopsis 63d16c4
- **swaggger:** add `GET` verb for `/healthz` 5a0bb54
- fix instruction to run mage's `build` target d35f513
- **api/handlers/liveness:** `Initialize()` function a8993b7
- **api/handlers/liveness:** `New()` function add5240
- **api/handlers/liveness:** `Handler` struct 06f7e41
- **api/handlers/liveness:** package constants 0e5f809
- **api/handlers/liveness:** `Status` enum 2ac0c5e
- **api/handlers/liveness:** package synopsis efe22ac
- **changelog:** updated changelog for v0.0.1 533f465

### Refactors

- **api/handlers/liveness:** update `handler()` to write response json without
  adding additional fields 985c2cf
- **api/core/router:** error propagation for http request handler init 02ce2eb
- **api/handlers/liveness:** error propagation for `New()` function f52d608

### Build System

- **go-releaser:** do not sign archive b931100
- **docker:** `gitpod` target 006af3c
- **just/bootstrap:** add `rust` toolchain 6d9e725
- **docker:** laxer package versioning in dockerfile 9db2bb8
- **docker:** fixed minor issues with `docker-bake.hcl` file 747b4f4
- **just/bootstrap:** add `bootstrap-os-pkgs` dependency to `bootstrap` target
  c397414
- **just/go:** fix `mage` install in `_build-go` target 7a6a5d9
- **just/commit:** fix `bootstrap-pre-commit` hook installation 495f420
- **just/semver:** changed changelog commit message pattern 9a5964e

### CI

- **docker:** detect `TAG` 5ca9bde
- **go:** add `go-releaser` step 74090a2
- **release:** fix deprecated flags in `go-releaser` step 0c641be
- update min go version to `1.20` 385d829
- **go:** fixed arguments in `test` step 2e45224
- **go:** fixed arguments in `build` step 55e6ae7
- **release:** remove `unshallow` step 2011e4f
- **release:** update `setup-go` and `goreleaser-action` actions ea4cc92

### Chore

- **gitpod:** source `gitpod` variables 168e63f
- **gitpod:** `rust` toolchain configuration e3e3d0d
- **gitpod:** add `bootstrap` task d2ce4f1
- **gitpod:** added more tools to the dockerfile 79f5f6d
- minor fix f9bc049
- minor fix 9ae11b2
- **pre-commit:** ensure Justfile is identical to upstream in `just-fmt.sh`
  script ee99516
- **gitpod:** replace `yarn` with `npm` in Dockerfile 927e64b
- change image in `.gitpod.yml` f494de7
- fix `.gitpod.yml` 8c3577e

### v0.0.1 (2023-02-12)

#### Features

- **api:** init 3974f95
- **api/core:** init 891fa92
- **api/errors:** init dc9c0cc
- **api/handlers:** init a9354fe
- **api/middlewares:** init 03877c6
- **api/registry:** init c093b2a
- **cmd/podinfo:** init a3c1ebf
- **internal/cryptoutil:** init 34a4597
- **internal/files:** init 4ac385c
- **internal/golang-lru:** init f1264aa
- **internal/locksutil:** init 09e59d1
- **internal/logger:** init e066d10
- **internal/multierror:** init a39b7c2
- **internal/pathmanager:** init 489996c
- **internal/permitpool:** init a29c5d7
- **internal/prettyprint:** init 2bc4ce1
- **internal/primitives:** init 1a67ad6
- **internal/runtimex:** init 76eb234
- **internal/testutils:** init c188be8
- **internal/urandom:** init 58ed4ff
- **internal/version:** init 4a98993
- **sdk/physical:** init f8d487e
- **internal/backoff/constant:** init f70b48d
- **internal/backoff/exponential:** init 437e93c
- **internal/cli/data:** init 9d115ad
- **internal/cli/decoder:** init 5b380a7
- **internal/cli/flagset:** init a233141
- **internal/cli/value:** init aef98ca
- **internal/golang-lru/simplelru:** init 8196569
- **internal/logger/slack:** init 1e97608
- **internal/radix-tree/immutable:** init 5dbb2e9
- **internal/radix-tree/mutable:** init fade8c6
- **sdk/api/address:** init 69918d7
- **sdk/api/fastjson:** init e5403e6
- **sdk/api/metrics:** init fd5e5e2
- **sdk/api/port:** init 979298e
- **sdk/api/proto:** init 3ddbd38
- **sdk/api/response:** init 76ab9c8
- **sdk/api/route:** init 2a7f6c7
- **sdk/physical/access:** init 829e8ef
- **sdk/physical/cache:** init 48e7940
- **sdk/physical/chroot:** init 5eb4b18
- **sdk/physical/encoding:** init b1895bf
- **sdk/physical/error-injector:** init fa52696
- **sdk/physical/latency:** init 7d4928f
- **sdk/physical/retry:** init 8660a4d
- **cmd/podinfo/commands/server:** init 540b954
- **cmd/podinfo/commands/version:** init df8d8d1
- **docker/release:** minimal `Dockerfile` for the binary c4cd6da

#### Documentation

- init 93e80f2
- **build/go:** `Build` target synopsis a9e2bfb
- **just/podinfo:** `run` target synopsis 966607f
- **just/podinfo:** `kill` target synopsis d4bb819
- **just/go:** `build-go` target synopsis 89c2bae
- **just/go:** `clean-go` target synopsis 7089f0b
- **just/go:** `lint-go` target synopsis 25b8479
- **just/go:** `format-go` target synopsis 1c62db8
- **just/go:** `bootstrap-go` target synopsis e37bcd0
- **just/go:** `_build` target synopsis 50231ce
- **just/go:** `_go` target synopsis f240c28
- **just/go:** `_lint-go` target synopsis 5ac43ba
- **just/semver:** `bootstrap-semver` target synopsis 92ec6b4
- **just/markdown:** `_format-markdown` target synopsis 2721074
- **just/justfile:** `format-just` target synopsis 78a7965
- **just/json:** `format-json` target synopsis 8c5778c
- **just/json:** `_format-json` target synopsis 6a417bd
- **just/bash:** `lint-bash` target synopsis 2e74c8d
- **just/bash:** `format-bash` target synopsis b59a5d7
- **just/bash:** `_lint-bash` target synopsis 08b1ac6
- **just/bash:** `_format-bash` target synopsis ef340bb
- **just/misc:** `snapshot` target synopsis 4026b9e
- **just/misc:** `vscode-tasks` target synopsis deb5e1d
- **just/semver:** `generate-changelog` target synopsis aa5683b
- **just/semver:** `patch-release` target synopsis 74afa35
- **just/semver:** `minor-release` target synopsis 2bde4d8
- **just/semver:** `major-release` target synopsis 03bb13d
- **just/semver:** variable synopsis a5db8af
- **just/commit:** `commit` target synopsis d02b807
- **just/commit:** `bootstrap-pre-commit` target synopsis c367eae
- **just/commit:** `_pre-commit` target synopsis 6aee1b8
- **just/git:** `git-add` target synopsis f84a9eb
- **just/git:** `_git-delta` target synopsis 3781d54
- **just/git:** `_git-delta` target synopsis b8d961d
- **just/bootstrap:** `_update-rust` target synopsis 0d1588e
- **just/bootstrap:** `_validate-rust` target synopsis b223255
- **just/bootstrap:** `_install-nodejs-package` target synopsis 58aabd6
- **just/bootstrap:** `_bootstrap-nodejs` target synopsis 00df92c
- **just/bootstrap:** `_core-pkgs` target synopsis 9a78d0a
- **just/bootstrap:** `_install-os-package` target synopsis b817954
- **just/bootstrap:** `_update-os-pkgs` target synopsis 4b94227
- **just/bootstrap:** `kary-comments` target synopsis f80992f
- **just:** common variables 9cc3843
- **build/go/targets/test:** `Target` function synopsis 7a17aad
- **build/go/targets:** package synopsis 40e593d
- **build/go/targets/test:** package synopsis 7e5375b
- **build/go/targets/build:** package synopsis 960062d
- **build/go/git:** package synopsis 2a1ffd4
- **build/go/version:** package synopsis 265da8c
- **api:** synopsis 295aad0
- **sdk:** synopsis 9f0d098
- **api/core:** synopsis 9ef7b34
- **api/handlers:** synopsis b0e0874
- **api/middlewares:** synopsis d3233c8
- **api/registry:** synopsis e801ead
- **internal/files:** synopsis a163f57
- **internal/golang-lru:** synopsis d63c0a7
- **internal/logger:** synopsis 5d65330
- **internal/permitpool:** synopsis 41feade
- **internal/prettyprint:** synopsis 87c51ba
- **internal/primitives:** synopsis 8a083ac
- **internal/runtimex:** synopsis dffbbdc
- **internal/testutils:** synopsis 3895ee7
- **internal/urandom:** synopsis ab5c0b9
- **sdk/api:** synopsis f657f44
- **sdk/physical:** synopsis 9706c76
- **cmd/podinfo/commands:** synopsis 108d9eb
- **internal/cli/data:** synopsis 3574a9c
- **internal/cli/decoder:** synopsis a13ad40
- **internal/cli/flagset:** synopsis 0a9d0c1
- **internal/logger/slack:** synopsis 83a078c
- **internal/radix-tree/immutable:** synopsis 78abe82
- **internal/radix-tree/mutable:** synopsis 4e53b91
- **sdk/api/address:** synopsis 23ab44b
- **sdk/api/fastjson:** synopsis 2a7c790
- **sdk/api/metrics:** synopsis c505728
- **sdk/api/port:** synopsis 1090e11
- **sdk/api/proto:** synopsis 197fad7
- **sdk/api/response:** synopsis 5cdbfab
- **sdk/api/route:** synopsis 6d23ed6
- **sdk/physical/access:** synopsis cea7bee
- **sdk/physical/cache:** synopsis 13e7f7d
- **sdk/physical/chroot:** synopsis 4b31bb4
- **sdk/physical/encoding:** synopsis 49a1e33
- **sdk/physical/error-injector:** synopsis 70446b1
- **sdk/physical/latency:** synopsis 1f906c5
- **sdk/physical/retry:** synopsis 13983aa
- **cmd/podinfo/commands/server:** synopsis 942978f
- **cmd/podinfo/commands/version:** synopsis 26ea212
- **docker-bake:** usage guide 5344b78
- **docker-bake:** `release` target synopsis e55beca
- **docker-bake:** `default` group synopsis baf2270
- **docker-bake:** `TAG` variable synopsis abe018f
- **docker-bake:** `AMD64` variable synopsis 7909900
- **docker-bake:** `ARM64` variable synopsis 750e36d
- **docker-bake:** `REGISTRY_USERNAME` variable synopsis 4c68a04
- **docker-bake:** `REGISTRY_HOSTNAME` variable synopsis 83e5534
- **docker-bake:** `LOCAL` variable synopsis e238d0e
- **github:** `programming_task` issue template dba5a78
- **github:** `feature_request` issue template 4124f12
- **github:** `bug_report` issue template c7d8513
- **github:** `api_endpoint_spec` issue template d60a2a9

#### Refactors

- **mage:** remove old files eabe9ea

#### Tests

- **api/core:** init 0598dda
- **internal/files:** init ffb9e6b
- **internal/logger:** init 1d6b435
- **internal/primitives:** init b3ebabc
- **internal/backoff/constant:** init c3255c2
- **internal/backoff/exponential:** init 588c83a
- **internal/cli/decoder:** init 1b3cfbe
- **internal/cli/value:** init 4aeaebb
- **sdk/api/fastjson:** init 01cd7be

#### Build System

- **just/core:** aggregator targets 3382f77
- **just/podinfo:** `run` target implementation e262bfa
- **just/podinfo:** `kill` target implementation 40a0637
- **just/go:** `build-go` target implementation d8670aa
- **just/go:** `clean-go` target implementation 8dc7e7c
- **just/go:** `lint-go` target implementation 14209fd
- **just/go:** `format-go` target implementation 0126826
- **just/go:** `bootstrap-go` target implementation 39b68ce
- **just/go:** `_build` hidden target implementation 47dc4e8
- **just/go:** `_go` hidden target implementation 26a247a
- **just/go:** `_lint-go` hidden target implementation 351a32a
- **just/semver:** `bootstrap-semver` target implementation 3ca58f4
- **just/markdown:** `_format-markdown` hidden target implementation 852e53e
- **just/justfile:** `format-just` target implementation eb8efa2
- **just/json:** `format-json` target implementation e540c34
- **just/json:** `_format-json` hidden target implementation 2754846
- **just/bash:** `lint-bash` target implementation 5cfd1f9
- **just/bash:** `format-bash` target implementation 05d97b0
- **just/bash:** `_lint-bash` hidden target implementation 2e56839
- **just/bash:** `_format-bash` hidden target implementation 2214242
- **just/misc:** `snapshot` target implementation 184f86e
- **just/misc:** `vscode-tasks` target implementation 8669e48
- **just/semver:** `generate-changelog` target implementation d9813b6
- **just/semver:** `patch-release` target implementation 0ee99d8
- **just/semver:** `minor-release` target implementation 47c8d9d
- **just/semver:** `major-release` target implementation 4ca9c5c
- **just/semver:** variable declaration fb9c78a
- **just/commit:** `commit` target implementation f574bf3
- **just/commit:** `bootstrap-pre-commit` target implementation 1243472
- **just/git:** `_pre-commit` hidden target implementation e5c5206
- **just/git:** `git-add` target implementation b053a64
- **just/git:** `git-fetch` target implementation 28764fb
- **just/git:** `_git-delta` hidden target implementation 33b956f
- **just/bootstrap:** `_install-rust-package` hidden target implementation
  81aab57
- **just/bootstrap:** `_update-rust` hidden target implementation 9b5d2d8
- **just/bootstrap:** `_validate-rust` hidden target implementation d8c6c38
- **just/bootstrap:** `_install-nodejs-package` hidden target implementation
  2a68b48
- **just/bootstrap:** `_bootstrap-nodejs` hidden target implementation 07c21bb
- **just/bootstrap:** `_core-pkgs` hidden target implementation 0cb70f3
- **just/bootstrap:** `_install-os-package` hidden target implementation
  9c70c26
- **just/bootstrap:** `_update-os-pkgs` hidden target implementation 7fb7d14
- **just/bootstrap:** `kary-comments` target implementation 8ffe083
- **just:** common config 86033e0
- **go:** moved from `magefile.go` 52e7886
- **go/targets/test:** moved from `mage/test` 84ef4a7
- **go/targets/build:** moved from `mage/build` a07a063
- **go/git:** moved from `mage/git` ae05d7a
- **go/version:** moved from `internal/version` 1906bdc
- **docker-bake:** `release` target implementation 778b6f7
- **docker-bake:** `TAG` variable declaration dd58811
- **docker-bake:** `AMD64` variable declaration 763dcb7
- **docker-bake:** `ARM64` variable declaration 7feeae8
- **docker-bake:** `REGISTRY_USERNAME` variable declaration c0493f8
- **docker-bake:** `REGISTRY_HOSTNAME` variable declaration 0cebd31
- **docker-bake:** `LOCAL` variable declaration a89e0be
- **docker/release:** added buildx installer/setup to the script 607964a
- **docker/release:** multi-arch builder script 2d320e6
- `go-releaser` config 25bea21
- **mage:** `magefile` init 2e7af35
- **test:** `test` mage targets 84e27ca
- **build:** `build` mage targets 5471c34
- **git:** `git` auxiliary library 4908743

#### CI

- **github:** `release` workflow fc1e8f4
- **github:** `go` workflow 299d59e
- **github:** `docker` workflow 4e9b676

#### Chore

- **pre-commit:** `just-fmt` pre-commit hook 57b16f2
- **pre-commit:** `md-fmt` pre-commit hook 8c4b6c4
- **docs:** fix `markdown` file exception `.gitignore` b501fa4
- **build:** local `.gitignore` 755f5d2
- **internal:** local `.gitignore` 825e7c2
- **sdk:** local `.gitignore` f55402b
- **cmd:** local `.gitignore` 6ff23fe
- **api:** local `.gitignore` a160290
- **docs:** local `.gitignore` fbfd5b7
- **assets:** local `.gitignore` 502e7a5
- **fixtures:** local `.gitignore` 7a098f4
- simplify `.editorconfig` file b61d49c
- fixed `.editorconfig` file a8db8f7
- **pre-commit:** `go-mod-tidy` pre-commit hook 713d5ab
- **pre-commit:** `go-fmt` pre-commit hook 2fbb351
- **github:** `CODEOWNERS` file 7d5618d
- **vscode:** `tasks` configuration file 8d9e1cd
- **vscode:** `settings` configuration file 3ec3faf
- **vscode:** `launch` configuration file 7ab2d30
- **vscode:** `extensions` configuration file 2f05faf
- `gitpod` config 6aeea05
- go vendor dependency management with `tools.go` 82fea78
- `.env` file init 24c7238
- `pre-commit` config file 3d8d786
- **linter:** `cspell` golang dictionary 7b20441
- **linter:** `cspell` generic dictionary b95cfbe
- **linter:** `cspell` config file 63d43ce
- **linter:** `golangci` config file 87a9658
- **linter:** `revive` config file 40abb30
- **linter:** `.markdownlintignore` file eb1bae8
- `.versionrc` file bee9668
- `.stignore` file 5bb284d
- `.editorconfig` file 640d343
- `.dockerignore` file aabfa50
- `gitignore` file 8c88606
- **linter:** `commitlint` config 4f34502

[1]: https://github.com/da-moon/northern-labs-interview/compare/v0.9.0...v1.0.0
[2]: https://github.com/da-moon/northern-labs-interview/compare/v0.8.2...v0.9.0
[3]: https://github.com/da-moon/northern-labs-interview/compare/v0.8.1...v0.8.2
[4]: https://github.com/da-moon/northern-labs-interview/compare/v0.8.0...v0.8.1
[5]: https://github.com/da-moon/northern-labs-interview/compare/v0.7.0...v0.8.0
[6]: https://github.com/da-moon/northern-labs-interview/compare/v0.6.3...v0.7.0
[7]: https://github.com/da-moon/northern-labs-interview/compare/v0.6.2...v0.6.3
[8]: https://github.com/da-moon/northern-labs-interview/compare/v0.6.1...v0.6.2
[9]: https://github.com/da-moon/northern-labs-interview/compare/v0.6.0...v0.6.1
[10]:
  https://github.com/da-moon/northern-labs-interview/compare/v0.5.0...v0.6.0
[11]:
  https://github.com/da-moon/northern-labs-interview/compare/v0.4.1...v0.5.0
[12]:
  https://github.com/da-moon/northern-labs-interview/compare/v0.4.0...v0.4.1
[13]:
  https://github.com/da-moon/northern-labs-interview/compare/v0.3.0...v0.4.0
[14]:
  https://github.com/da-moon/northern-labs-interview/compare/v0.2.2...v0.3.0
[15]:
  https://github.com/da-moon/northern-labs-interview/compare/v0.2.1...v0.2.2
[16]:
  https://github.com/da-moon/northern-labs-interview/compare/v0.2.0...v0.2.1
[17]:
  https://github.com/da-moon/northern-labs-interview/compare/v0.1.1...v0.2.0
[18]:
  https://github.com/da-moon/northern-labs-interview/compare/v0.1.0...v0.1.1
[19]:
  https://github.com/da-moon/northern-labs-interview/compare/v0.0.1...v0.1.0
