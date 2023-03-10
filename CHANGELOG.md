# Changelog

### [v1.0.1][1] (2023-02-21)

#### Documentation

- **changelog:** v1.0.0 ([eb5e9f3][2])

#### Chore

- package rename ([3a129a3][3])
- **docker-compose:** add `podinfo` service ([32a7e7d][4])

## [v1.0.0][5] (2023-02-15)

### Documentation

- **changelog:** v0.9.0 ([8ead18d][6])

### Build System

- **just/podinfo:** `/cache/` api group recipes ([2108ed4][7])

## [v0.9.0][8] (2023-02-15)

### Bug Fixes

- **api/handlers/cache:** all cache endpoints are working ([63e63c7][9])
- **api/handlers/cache:** enabled `/cache` api group ([718eae8][10])

### Documentation

- **changelog:** v0.9.0 ([9f50cc4][11])

### Refactors

- better logging ([f61c835][12])
- **build/go/version:** remove unnecessary information ([6999c69][13])

### Chore

- **docker-compose:** fix redis listener ([3ad11d9][14])
- **docker-compose:** `redis` service configuration file ([d3178be][15])
- **docker-compose:** `redis` service ([5f696a1][16])

### [v0.8.2][17] (2023-02-15)

#### Documentation

- **changelog:** v0.9.0 ([868a12d][18])

#### CI

- **docker:** checking to see if docker credenials exist ([2e7d3fa][19])
- **docker:** checking to see if docker credenials exist ([06a0f86][20])
- **docker:** trigger build after adding secrets ([e8be68d][21])
- **docker:** fix `buildx` id collision ([6f1f2a8][22])

### [v0.8.1][23] (2023-02-15)

#### Documentation

- **changelog:** v0.8.0 ([9522755][24])

#### CI

- **dockeer:** trigger build on changes to any `*.go` file ([7ec1748][25])

## [v0.8.0][26] (2023-02-15)

### Documentation

- add some instructions and context ([344ff9a][27])
- update `just` target in main `README` file ([68e9779][28])
- add `Roadmap` ([a3b0242][29])
- `swagger` yaml ([efb738c][30])
- **changelog:** v0.7.0 ([13a5032][31])

## [v0.7.0][32] (2023-02-15)

### Features

- **cmd/podinfo/commands/server:** parse redis info with flags ([8e57115][33])
- **api/handlers/cache:** register **DELETE** `/cache/{key}` route in API group
  `Register()` ([a5ab3a8][34])
- **api/handlers/cache:** register **DELETE** `/cache/{key}` route in API group
  `Register()` ([d12f7db][35])
- **api/handlers/cache/delete:** `handler` implementation ([07c94d5][36])

### Documentation

- added code stats ([3e0e1ef][37])
- **api/handlers/cache/delete:** package synopsis ([d0e25cb][38])
- **changelog:** v0.7.0 ([8b5e5d2][39])

### [v0.6.3][40] (2023-02-15)

#### Features

- **api/handlers/cache:** register **GET** `/cache/{key}` route in API group
  `Register()` ([ab70334][41])
- **api/handlers/cache/get:** route `register` implementation ([9935643][42])
- **api/handlers/cache/get:** `handler` implementation ([f154df8][43])

#### Documentation

- **api/handlers/cache/doc:** package synopsis ([9de9d35][44])
- **changelog:** v0.7.0 ([61739b5][45])

### [v0.6.2][46] (2023-02-15)

#### Features

- **api/handlers/cache:** register **PUT** `/cache/{key}` route in API group
  `Register()` ([b7cb2ba][47])
- **api/handlers/cache/put:** route `register` implementation ([6696e26][48])
- **api/handlers/cache/put:** `handler` implementation ([15a177c][49])

#### Documentation

- **api/handlers/cache/put:** package synopsis ([26257c2][50])
- **changelog:** v0.7.0 ([1b6e7f1][51])

### [v0.6.1][52] (2023-02-15)

#### Features

- **api/handlers/cache:** register **POST** `/cache/{key}` route in API group
  `Register()` ([f2f2164][53])
- **api/handlers/cache/post:** route `register` implementation ([fb15f04][54])
- **api/handlers/cache/post:** `handler` implementation ([86f9e9d][55])

#### Documentation

- **api/handlers/cache:** package synopsis ([1e8dc7e][56])
- **api/handlers/cache/post:** package synopsis ([f878c09][57])
- **changelog:** v0.6.0 ([cf372a5][58])

## [v0.6.0][59] (2023-02-15)

### Features

- **api/handlers/cache/shared:** `preflight` request handler ([489db81][60])
- **api/handlers/cache/shared:** `API group` shared constructs ([47b687d][61])
- **cmd/podinfo/commands/server:** flags to read `redis` config ([76a07eb][62])
- **api/core:** add `redis` configuration fields and methods ([cc54fcf][63])
- **internal/cli/value:** handle `int` type cli flags ([d26c406][64])

### Documentation

- **api/handlers/cache/shared:** package synopsis ([2732306][65])
- **changelog:** v0.5.0 ([c313777][66])

### Refactors

- **internal/cli/flagset:** more flexibilty in prettifying generated flag usage
  ([893a1e2][67])
- **internal/cli/flagset:** more flexiblity in self-documenting flags
  ([fc63a46][68])

### Build System

- **docker/podinfo:** minor changes to build script ([9e7f04a][69])

### Chore

- **re-commit:** hook updates ([bff048a][70])

## [v0.5.0][71] (2023-02-15)

### [v0.4.1][72] (2023-02-15)

#### Features

- **api/handlers:** register **GET** `/delay/{seconds}` route in `Initialize()`
  ([5a9dc8e][73])
- **api/handlers/delay:** route `register` implementation ([f39830b][74])
- **api/handlers/delay:** `handler` implementation ([bf469a3][75])
- **api/handlers/delay:** `response` setup ([d22dc78][76])

#### Documentation

- **changelog:** v0.5.0 ([c313777][66])
- **api/handlers/delay:** package synopsis ([e260245][77])
- **changelog:** v0.4.0 ([d7d5af7][78])

#### Build System

- **build/just/podinfo:** `delay-probe` recipe ([9d0b55f][79])

## [v0.4.0][80] (2023-02-15)

### Features

- **api/handlers:** register `/headers` route in `Initialize()` ([10cef5d][81])
- **api/handlers/headers:** route `register` implementation ([989f948][82])
- **api/handlers/headers:** `handler` implementation ([9871e64][83])
- **api/handlers/headers:** `response` type ([019355e][84])

### Documentation

- **api/handlers/headers:** package synopsis ([489218a][85])

### Build System

- **build/just/podinfo:** `headers-probe` recipe ([239a061][86])

## [v0.3.0][87] (2023-02-15)

### Features

- **api/handlers:** register `/env` route in `Initialize()` ([2d98b6a][88])
- **api/handlers/env:** routte `register` implementation ([958bc47][89])
- **api/handlers/env:** `handler` implementation ([0ef5a8f][90])
- **api/handlers/env:** `response` type ([877b7e4][91])

### Documentation

- **api/handlers/env:** package synopsis ([edc489e][92])

### Build System

- **build/just/podinfo:** `env-probe` recipe ([b66f28b][93])

### [v0.2.2][94] (2023-02-15)

#### Features

- **api/handlers:** register `/readyz/enable` route in `Initialize()`
  ([dead392][95])
- **api/handlers/readiness/enable:** route init ([b56a400][96])
- **api/handlers/readiness/enable:** handler setup ([9356de0][97])

#### Documentation

- **api/handlers/readiness/enable:** package synopsis ([faeaf50][98])

#### Build System

- **build/just/podinfo:** `readiness-probe-enable` recipe ([65c6272][99])

### [v0.2.1][100] (2023-02-15)

#### Features

- **api/handlers:** register the route in `Initialize()` ([ae99b84][101])
- **api/handlers/readiness/disable:** `handler.Register()` implementation
  ([a3851e6][102])
- **api/handlers/readiness/disable:** `HandlerFn` implementation
  ([6297492][103])
- **api/handlers/readiness/disable:** `handler.GetLogger()` synopsis
  ([3e7a28f][104])
- **api/handlers/readiness/disable:** `handler.GetLogger()` implementation
  ([fcdb30b][105])
- **api/handlers/readiness/disable:** `handler.SetLogger()` synopsis
  ([ce5be1d][106])
- **api/handlers/readiness/disable:** `handler.SetLogger()` implementation
  ([b39a633][107])
- **api/handlers/readiness/disable:** `handler` struct implementation
  ([32376d6][108])

#### Documentation

- **api/handlers/readiness/disable:** `HandlerFn` synopsis ([ac1c0ee][109])
- **api/handlers/readiness/disable:** `handler` struct synopsis
  ([5e8ae1e][110])
- **api/handlers/readiness/disable:** package synopsis ([4753ee2][111])
- **changelog:** v0.2.0 ([2ec37fc][112])

#### Build System

- **build/just/podinfo:** `readiness-probe-disable` recipe ([f0efeb5][113])

## [v0.2.0][114] (2023-02-15)

### Features

- **api/handlers/readiness:** `handler.Register()` implementation
  ([229f6ef][115])
- **api/handlers/readiness:** `init()` implementation ([316c88d][116])
- **api/handlers/readiness:** `HandlerFn` implementation ([fa4e4e0][117])
- **api/handlers/readiness:** `handler.GetStatus()` implementation
  ([6975ceb][118])
- **api/handlers/readiness:** `handler.SetStatus()` implementation
  ([0ae4c75][119])
- **api/handlers/readiness:** `handler.GetLogger()` synopsis ([1fb4903][120])
- **api/handlers/readiness:** `handler.GetLogger()` implementation
  ([2af2f25][121])
- **api/handlers/readiness:** `handler.SetLogger()` implementation
  ([a34f22d][122])
- **api/handlers/readiness:** `handler` struct implementation ([4a73558][123])
- **api/handlers/readiness:** `Response` struct implementation ([8ea8855][124])
- **api/handlers/readiness:** `Status.String()` implementation ([f81e205][125])
- **api/handlers/readiness:** `Status` enum ([c45a277][126])
- **sdk/api/response:** `LogEntry()` implementation ([c62180d][127])

### Documentation

- **build/just/podinfo:** `readiness-probe` recipe ([a989ce1][128])
- **api/handlers/readiness:** `HandlerFn` synopsis ([5b226c3][129])
- **api/handlers/readiness:** `handler.GetStatus()` synopsis ([7f85163][130])
- **api/handlers/readiness:** `handler.SetStatus()` synopsis ([db54578][131])
- **api/handlers/readiness:** `handler.SetLogger()` synopsis ([2ac7d67][132])
- **api/handlers/readiness:** `handler` struct synopsis ([14fcb57][133])
- **api/handlers/readiness:** `Response` struct synopsis ([7c7ecff][134])
- **api/handlers/readiness:** `Status.String()` synopsis ([777320c][135])
- **api/handlers/readiness:** `Status` enum ([9788dfb][136])
- **api/handlers/readiness:** package synopsis ([92c17f3][137])
- **sdk/api/response:** `LogEntry()` synopsis ([c9c5f7d][138])
- **changelog:** v0.2.0 ([c46195b][139])

### Refactors

- **api/handlers/liveness:** new pattern of writing `HandlerFn`
  ([d270881][140])
- **sdk/api/response:** `Write` function ([66fd066][141])
- **api/handlers/liveness:** change route impl pattern ([6efbf25][142])

### Build System

- **build/just/podinfo:** `readiness-probe` recipe ([14a6fe0][143])

### Chore

- **gitpod:** fix `gitpod` shell config in dockerfile ([ad05fc9][144])

### [v0.1.1][145] (2023-02-13)

#### Documentation

- **changelog:** `v0.1.0` ([e1bdb33][146])

#### Build System

- **just/git:** prune refs in `git-fetch` recipe ([69d5e05][147])

## [v0.1.0][148] (2023-02-13)

### Features

- **sdk/api/response:** `WriteJSON()` function implementation ([ac8c96c][149])
- **sdk/api/response:** `WriteSuccessfulJSONRaw()` function implementation
  ([f264be1][150])
- **api/handlers:** remove `Prefix` from pre-flight handler ([ae48e9d][151])
- **api/handlers:** Register and initialize `kubernetes-liveness-probe` route
  ([1c416f5][152])
- **api/handlers/liveness:** `Initialize()` function ([73548a0][153])
- **api/handlers/liveness:** `New()` function ([016f07c][154])
- **api/handlers/liveness:** `Handler` struct ([133a5a8][155])
- **api/handlers/liveness:** package constants ([a0b8f71][156])
- **api/handlers/liveness:** `Response` struct ([9bc1341][157])

### Bug Fixes

- **sdk/api/response:** address linter complaints ([a9eed4e][158])

### Documentation

- **docker/buildx:** `gitpod` target ([11b8055][159])
- `liveness-probe` instruction ([fc5852c][160])
- **sdk/api/response:** `WriteJSON()` function synopsis ([7e25e0a][161])
- **sdk/api/response:** `WriteSuccessfulJSONRaw()` synopsis ([63d16c4][162])
- **swaggger:** add `GET` verb for `/healthz` ([5a0bb54][163])
- fix instruction to run mage's `build` target ([d35f513][164])
- **api/handlers/liveness:** `Initialize()` function ([a8993b7][165])
- **api/handlers/liveness:** `New()` function ([add5240][166])
- **api/handlers/liveness:** `Handler` struct ([06f7e41][167])
- **api/handlers/liveness:** package constants ([0e5f809][168])
- **api/handlers/liveness:** `Status` enum ([2ac0c5e][169])
- **api/handlers/liveness:** package synopsis ([efe22ac][170])
- **changelog:** updated changelog for v0.0.1 ([533f465][171])

### Refactors

- **api/handlers/liveness:** update `handler()` to write response json without
  adding additional fields ([985c2cf][172])
- **api/core/router:** error propagation for http request handler init
  ([02ce2eb][173])
- **api/handlers/liveness:** error propagation for `New()` function
  ([f52d608][174])

### Build System

- **go-releaser:** do not sign archive ([b931100][175])
- **docker:** `gitpod` target ([006af3c][176])
- **just/bootstrap:** add `rust` toolchain ([6d9e725][177])
- **docker:** laxer package versioning in dockerfile ([9db2bb8][178])
- **docker:** fixed minor issues with `docker-bake.hcl` file ([747b4f4][179])
- **just/bootstrap:** add `bootstrap-os-pkgs` dependency to `bootstrap` target
  ([c397414][180])
- **just/go:** fix `mage` install in `_build-go` target ([7a6a5d9][181])
- **just/commit:** fix `bootstrap-pre-commit` hook installation
  ([495f420][182])
- **just/semver:** changed changelog commit message pattern ([9a5964e][183])

### CI

- **docker:** detect `TAG` ([5ca9bde][184])
- **go:** add `go-releaser` step ([74090a2][185])
- **release:** fix deprecated flags in `go-releaser` step ([0c641be][186])
- update min go version to `1.20` ([385d829][187])
- **go:** fixed arguments in `test` step ([2e45224][188])
- **go:** fixed arguments in `build` step ([55e6ae7][189])
- **release:** remove `unshallow` step ([2011e4f][190])
- **release:** update `setup-go` and `goreleaser-action` actions
  ([ea4cc92][191])

### Chore

- **gitpod:** source `gitpod` variables ([168e63f][192])
- **gitpod:** `rust` toolchain configuration ([e3e3d0d][193])
- **gitpod:** add `bootstrap` task ([d2ce4f1][194])
- **gitpod:** added more tools to the dockerfile ([79f5f6d][195])
- minor fix ([f9bc049][196])
- minor fix ([9ae11b2][197])
- **pre-commit:** ensure Justfile is identical to upstream in `just-fmt.sh`
  script ([ee99516][198])
- **gitpod:** replace `yarn` with `npm` in Dockerfile ([927e64b][199])
- change image in `.gitpod.yml` ([f494de7][200])
- fix `.gitpod.yml` ([8c3577e][201])

### v0.0.1 (2023-02-12)

#### Features

- **api:** init ([3974f95][202])
- **api/core:** init ([891fa92][203])
- **api/errors:** init ([dc9c0cc][204])
- **api/handlers:** init ([a9354fe][205])
- **api/middlewares:** init ([03877c6][206])
- **api/registry:** init ([c093b2a][207])
- **cmd/podinfo:** init ([a3c1ebf][208])
- **internal/cryptoutil:** init ([34a4597][209])
- **internal/files:** init ([4ac385c][210])
- **internal/golang-lru:** init ([f1264aa][211])
- **internal/locksutil:** init ([09e59d1][212])
- **internal/logger:** init ([e066d10][213])
- **internal/multierror:** init ([a39b7c2][214])
- **internal/pathmanager:** init ([489996c][215])
- **internal/permitpool:** init ([a29c5d7][216])
- **internal/prettyprint:** init ([2bc4ce1][217])
- **internal/primitives:** init ([1a67ad6][218])
- **internal/runtimex:** init ([76eb234][219])
- **internal/testutils:** init ([c188be8][220])
- **internal/urandom:** init ([58ed4ff][221])
- **internal/version:** init ([4a98993][222])
- **sdk/physical:** init ([f8d487e][223])
- **internal/backoff/constant:** init ([f70b48d][224])
- **internal/backoff/exponential:** init ([437e93c][225])
- **internal/cli/data:** init ([9d115ad][226])
- **internal/cli/decoder:** init ([5b380a7][227])
- **internal/cli/flagset:** init ([a233141][228])
- **internal/cli/value:** init ([aef98ca][229])
- **internal/golang-lru/simplelru:** init ([8196569][230])
- **internal/logger/slack:** init ([1e97608][231])
- **internal/radix-tree/immutable:** init ([5dbb2e9][232])
- **internal/radix-tree/mutable:** init ([fade8c6][233])
- **sdk/api/address:** init ([69918d7][234])
- **sdk/api/fastjson:** init ([e5403e6][235])
- **sdk/api/metrics:** init ([fd5e5e2][236])
- **sdk/api/port:** init ([979298e][237])
- **sdk/api/proto:** init ([3ddbd38][238])
- **sdk/api/response:** init ([76ab9c8][239])
- **sdk/api/route:** init ([2a7f6c7][240])
- **sdk/physical/access:** init ([829e8ef][241])
- **sdk/physical/cache:** init ([48e7940][242])
- **sdk/physical/chroot:** init ([5eb4b18][243])
- **sdk/physical/encoding:** init ([b1895bf][244])
- **sdk/physical/error-injector:** init ([fa52696][245])
- **sdk/physical/latency:** init ([7d4928f][246])
- **sdk/physical/retry:** init ([8660a4d][247])
- **cmd/podinfo/commands/server:** init ([540b954][248])
- **cmd/podinfo/commands/version:** init ([df8d8d1][249])
- **docker/release:** minimal `Dockerfile` for the binary ([c4cd6da][250])

#### Documentation

- init ([93e80f2][251])
- **build/go:** `Build` target synopsis ([a9e2bfb][252])
- **just/podinfo:** `run` target synopsis ([966607f][253])
- **just/podinfo:** `kill` target synopsis ([d4bb819][254])
- **just/go:** `build-go` target synopsis ([89c2bae][255])
- **just/go:** `clean-go` target synopsis ([7089f0b][256])
- **just/go:** `lint-go` target synopsis ([25b8479][257])
- **just/go:** `format-go` target synopsis ([1c62db8][258])
- **just/go:** `bootstrap-go` target synopsis ([e37bcd0][259])
- **just/go:** `_build` target synopsis ([50231ce][260])
- **just/go:** `_go` target synopsis ([f240c28][261])
- **just/go:** `_lint-go` target synopsis ([5ac43ba][262])
- **just/semver:** `bootstrap-semver` target synopsis ([92ec6b4][263])
- **just/markdown:** `_format-markdown` target synopsis ([2721074][264])
- **just/justfile:** `format-just` target synopsis ([78a7965][265])
- **just/json:** `format-json` target synopsis ([8c5778c][266])
- **just/json:** `_format-json` target synopsis ([6a417bd][267])
- **just/bash:** `lint-bash` target synopsis ([2e74c8d][268])
- **just/bash:** `format-bash` target synopsis ([b59a5d7][269])
- **just/bash:** `_lint-bash` target synopsis ([08b1ac6][270])
- **just/bash:** `_format-bash` target synopsis ([ef340bb][271])
- **just/misc:** `snapshot` target synopsis ([4026b9e][272])
- **just/misc:** `vscode-tasks` target synopsis ([deb5e1d][273])
- **just/semver:** `generate-changelog` target synopsis ([aa5683b][274])
- **just/semver:** `patch-release` target synopsis ([74afa35][275])
- **just/semver:** `minor-release` target synopsis ([2bde4d8][276])
- **just/semver:** `major-release` target synopsis ([03bb13d][277])
- **just/semver:** variable synopsis ([a5db8af][278])
- **just/commit:** `commit` target synopsis ([d02b807][279])
- **just/commit:** `bootstrap-pre-commit` target synopsis ([c367eae][280])
- **just/commit:** `_pre-commit` target synopsis ([6aee1b8][281])
- **just/git:** `git-add` target synopsis ([f84a9eb][282])
- **just/git:** `_git-delta` target synopsis ([3781d54][283])
- **just/git:** `_git-delta` target synopsis ([b8d961d][284])
- **just/bootstrap:** `_update-rust` target synopsis ([0d1588e][285])
- **just/bootstrap:** `_validate-rust` target synopsis ([b223255][286])
- **just/bootstrap:** `_install-nodejs-package` target synopsis
  ([58aabd6][287])
- **just/bootstrap:** `_bootstrap-nodejs` target synopsis ([00df92c][288])
- **just/bootstrap:** `_core-pkgs` target synopsis ([9a78d0a][289])
- **just/bootstrap:** `_install-os-package` target synopsis ([b817954][290])
- **just/bootstrap:** `_update-os-pkgs` target synopsis ([4b94227][291])
- **just/bootstrap:** `kary-comments` target synopsis ([f80992f][292])
- **just:** common variables ([9cc3843][293])
- **build/go/targets/test:** `Target` function synopsis ([7a17aad][294])
- **build/go/targets:** package synopsis ([40e593d][295])
- **build/go/targets/test:** package synopsis ([7e5375b][296])
- **build/go/targets/build:** package synopsis ([960062d][297])
- **build/go/git:** package synopsis ([2a1ffd4][298])
- **build/go/version:** package synopsis ([265da8c][299])
- **api:** synopsis ([295aad0][300])
- **sdk:** synopsis ([9f0d098][301])
- **api/core:** synopsis ([9ef7b34][302])
- **api/handlers:** synopsis ([b0e0874][303])
- **api/middlewares:** synopsis ([d3233c8][304])
- **api/registry:** synopsis ([e801ead][305])
- **internal/files:** synopsis ([a163f57][306])
- **internal/golang-lru:** synopsis ([d63c0a7][307])
- **internal/logger:** synopsis ([5d65330][308])
- **internal/permitpool:** synopsis ([41feade][309])
- **internal/prettyprint:** synopsis ([87c51ba][310])
- **internal/primitives:** synopsis ([8a083ac][311])
- **internal/runtimex:** synopsis ([dffbbdc][312])
- **internal/testutils:** synopsis ([3895ee7][313])
- **internal/urandom:** synopsis ([ab5c0b9][314])
- **sdk/api:** synopsis ([f657f44][315])
- **sdk/physical:** synopsis ([9706c76][316])
- **cmd/podinfo/commands:** synopsis ([108d9eb][317])
- **internal/cli/data:** synopsis ([3574a9c][318])
- **internal/cli/decoder:** synopsis ([a13ad40][319])
- **internal/cli/flagset:** synopsis ([0a9d0c1][320])
- **internal/logger/slack:** synopsis ([83a078c][321])
- **internal/radix-tree/immutable:** synopsis ([78abe82][322])
- **internal/radix-tree/mutable:** synopsis ([4e53b91][323])
- **sdk/api/address:** synopsis ([23ab44b][324])
- **sdk/api/fastjson:** synopsis ([2a7c790][325])
- **sdk/api/metrics:** synopsis ([c505728][326])
- **sdk/api/port:** synopsis ([1090e11][327])
- **sdk/api/proto:** synopsis ([197fad7][328])
- **sdk/api/response:** synopsis ([5cdbfab][329])
- **sdk/api/route:** synopsis ([6d23ed6][330])
- **sdk/physical/access:** synopsis ([cea7bee][331])
- **sdk/physical/cache:** synopsis ([13e7f7d][332])
- **sdk/physical/chroot:** synopsis ([4b31bb4][333])
- **sdk/physical/encoding:** synopsis ([49a1e33][334])
- **sdk/physical/error-injector:** synopsis ([70446b1][335])
- **sdk/physical/latency:** synopsis ([1f906c5][336])
- **sdk/physical/retry:** synopsis ([13983aa][337])
- **cmd/podinfo/commands/server:** synopsis ([942978f][338])
- **cmd/podinfo/commands/version:** synopsis ([26ea212][339])
- **docker-bake:** usage guide ([5344b78][340])
- **docker-bake:** `release` target synopsis ([e55beca][341])
- **docker-bake:** `default` group synopsis ([baf2270][342])
- **docker-bake:** `TAG` variable synopsis ([abe018f][343])
- **docker-bake:** `AMD64` variable synopsis ([7909900][344])
- **docker-bake:** `ARM64` variable synopsis ([750e36d][345])
- **docker-bake:** `REGISTRY_USERNAME` variable synopsis ([4c68a04][346])
- **docker-bake:** `REGISTRY_HOSTNAME` variable synopsis ([83e5534][347])
- **docker-bake:** `LOCAL` variable synopsis ([e238d0e][348])
- **github:** `programming_task` issue template ([dba5a78][349])
- **github:** `feature_request` issue template ([4124f12][350])
- **github:** `bug_report` issue template ([c7d8513][351])
- **github:** `api_endpoint_spec` issue template ([d60a2a9][352])

#### Refactors

- **mage:** remove old files ([eabe9ea][353])

#### Tests

- **api/core:** init ([0598dda][354])
- **internal/files:** init ([ffb9e6b][355])
- **internal/logger:** init ([1d6b435][356])
- **internal/primitives:** init ([b3ebabc][357])
- **internal/backoff/constant:** init ([c3255c2][358])
- **internal/backoff/exponential:** init ([588c83a][359])
- **internal/cli/decoder:** init ([1b3cfbe][360])
- **internal/cli/value:** init ([4aeaebb][361])
- **sdk/api/fastjson:** init ([01cd7be][362])

#### Build System

- **just/core:** aggregator targets ([3382f77][363])
- **just/podinfo:** `run` target implementation ([e262bfa][364])
- **just/podinfo:** `kill` target implementation ([40a0637][365])
- **just/go:** `build-go` target implementation ([d8670aa][366])
- **just/go:** `clean-go` target implementation ([8dc7e7c][367])
- **just/go:** `lint-go` target implementation ([14209fd][368])
- **just/go:** `format-go` target implementation ([0126826][369])
- **just/go:** `bootstrap-go` target implementation ([39b68ce][370])
- **just/go:** `_build` hidden target implementation ([47dc4e8][371])
- **just/go:** `_go` hidden target implementation ([26a247a][372])
- **just/go:** `_lint-go` hidden target implementation ([351a32a][373])
- **just/semver:** `bootstrap-semver` target implementation ([3ca58f4][374])
- **just/markdown:** `_format-markdown` hidden target implementation
  ([852e53e][375])
- **just/justfile:** `format-just` target implementation ([eb8efa2][376])
- **just/json:** `format-json` target implementation ([e540c34][377])
- **just/json:** `_format-json` hidden target implementation ([2754846][378])
- **just/bash:** `lint-bash` target implementation ([5cfd1f9][379])
- **just/bash:** `format-bash` target implementation ([05d97b0][380])
- **just/bash:** `_lint-bash` hidden target implementation ([2e56839][381])
- **just/bash:** `_format-bash` hidden target implementation ([2214242][382])
- **just/misc:** `snapshot` target implementation ([184f86e][383])
- **just/misc:** `vscode-tasks` target implementation ([8669e48][384])
- **just/semver:** `generate-changelog` target implementation ([d9813b6][385])
- **just/semver:** `patch-release` target implementation ([0ee99d8][386])
- **just/semver:** `minor-release` target implementation ([47c8d9d][387])
- **just/semver:** `major-release` target implementation ([4ca9c5c][388])
- **just/semver:** variable declaration ([fb9c78a][389])
- **just/commit:** `commit` target implementation ([f574bf3][390])
- **just/commit:** `bootstrap-pre-commit` target implementation
  ([1243472][391])
- **just/git:** `_pre-commit` hidden target implementation ([e5c5206][392])
- **just/git:** `git-add` target implementation ([b053a64][393])
- **just/git:** `git-fetch` target implementation ([28764fb][394])
- **just/git:** `_git-delta` hidden target implementation ([33b956f][395])
- **just/bootstrap:** `_install-rust-package` hidden target implementation
  ([81aab57][396])
- **just/bootstrap:** `_update-rust` hidden target implementation
  ([9b5d2d8][397])
- **just/bootstrap:** `_validate-rust` hidden target implementation
  ([d8c6c38][398])
- **just/bootstrap:** `_install-nodejs-package` hidden target implementation
  ([2a68b48][399])
- **just/bootstrap:** `_bootstrap-nodejs` hidden target implementation
  ([07c21bb][400])
- **just/bootstrap:** `_core-pkgs` hidden target implementation
  ([0cb70f3][401])
- **just/bootstrap:** `_install-os-package` hidden target implementation
  ([9c70c26][402])
- **just/bootstrap:** `_update-os-pkgs` hidden target implementation
  ([7fb7d14][403])
- **just/bootstrap:** `kary-comments` target implementation ([8ffe083][404])
- **just:** common config ([86033e0][405])
- **go:** moved from `magefile.go` ([52e7886][406])
- **go/targets/test:** moved from `mage/test` ([84ef4a7][407])
- **go/targets/build:** moved from `mage/build` ([a07a063][408])
- **go/git:** moved from `mage/git` ([ae05d7a][409])
- **go/version:** moved from `internal/version` ([1906bdc][410])
- **docker-bake:** `release` target implementation ([778b6f7][411])
- **docker-bake:** `TAG` variable declaration ([dd58811][412])
- **docker-bake:** `AMD64` variable declaration ([763dcb7][413])
- **docker-bake:** `ARM64` variable declaration ([7feeae8][414])
- **docker-bake:** `REGISTRY_USERNAME` variable declaration ([c0493f8][415])
- **docker-bake:** `REGISTRY_HOSTNAME` variable declaration ([0cebd31][416])
- **docker-bake:** `LOCAL` variable declaration ([a89e0be][417])
- **docker/release:** added buildx installer/setup to the script
  ([607964a][418])
- **docker/release:** multi-arch builder script ([2d320e6][419])
- `go-releaser` config ([25bea21][420])
- **mage:** `magefile` init ([2e7af35][421])
- **test:** `test` mage targets ([84e27ca][422])
- **build:** `build` mage targets ([5471c34][423])
- **git:** `git` auxiliary library ([4908743][424])

#### CI

- **github:** `release` workflow ([fc1e8f4][425])
- **github:** `go` workflow ([299d59e][426])
- **github:** `docker` workflow ([4e9b676][427])

#### Chore

- **pre-commit:** `just-fmt` pre-commit hook ([57b16f2][428])
- **pre-commit:** `md-fmt` pre-commit hook ([8c4b6c4][429])
- **docs:** fix `markdown` file exception `.gitignore` ([b501fa4][430])
- **build:** local `.gitignore` ([755f5d2][431])
- **internal:** local `.gitignore` ([825e7c2][432])
- **sdk:** local `.gitignore` ([f55402b][433])
- **cmd:** local `.gitignore` ([6ff23fe][434])
- **api:** local `.gitignore` ([a160290][435])
- **docs:** local `.gitignore` ([fbfd5b7][436])
- **assets:** local `.gitignore` ([502e7a5][437])
- **fixtures:** local `.gitignore` ([7a098f4][438])
- simplify `.editorconfig` file ([b61d49c][439])
- fixed `.editorconfig` file ([a8db8f7][440])
- **pre-commit:** `go-mod-tidy` pre-commit hook ([713d5ab][441])
- **pre-commit:** `go-fmt` pre-commit hook ([2fbb351][442])
- **github:** `CODEOWNERS` file ([7d5618d][443])
- **vscode:** `tasks` configuration file ([8d9e1cd][444])
- **vscode:** `settings` configuration file ([3ec3faf][445])
- **vscode:** `launch` configuration file ([7ab2d30][446])
- **vscode:** `extensions` configuration file ([2f05faf][447])
- `gitpod` config ([6aeea05][448])
- go vendor dependency management with `tools.go` ([82fea78][449])
- `.env` file init ([24c7238][450])
- `pre-commit` config file ([3d8d786][451])
- **linter:** `cspell` golang dictionary ([7b20441][452])
- **linter:** `cspell` generic dictionary ([b95cfbe][453])
- **linter:** `cspell` config file ([63d43ce][454])
- **linter:** `golangci` config file ([87a9658][455])
- **linter:** `revive` config file ([40abb30][456])
- **linter:** `.markdownlintignore` file ([eb1bae8][457])
- `.versionrc` file ([bee9668][458])
- `.stignore` file ([5bb284d][459])
- `.editorconfig` file ([640d343][460])
- `.dockerignore` file ([aabfa50][461])
- `gitignore` file ([8c88606][462])
- **linter:** `commitlint` config ([4f34502][463])

[1]: https://github.com/da-moon/podinfo/compare/v1.0.0...v1.0.1
[2]:
  https://github.com/da-moon/podinfo/commit/eb5e9f3b79a95bf7f24bfe11b7abb056dafd0d2c
[3]:
  https://github.com/da-moon/podinfo/commit/3a129a3774a59321636aa1f097f9c70900996f9d
[4]:
  https://github.com/da-moon/podinfo/commit/32a7e7d9d74d0872834bc506ea93233bdcb33845
[5]: https://github.com/da-moon/podinfo/compare/v0.9.0...v1.0.0
[6]:
  https://github.com/da-moon/podinfo/commit/8ead18d61f80ef2ec763c807a58b440e70b25375
[7]:
  https://github.com/da-moon/podinfo/commit/2108ed484385d574acbba42ea428de971d007d3b
[8]: https://github.com/da-moon/podinfo/compare/v0.8.2...v0.9.0
[9]:
  https://github.com/da-moon/podinfo/commit/63e63c7610a860466f46e1d633ef49b44c31c081
[10]:
  https://github.com/da-moon/podinfo/commit/718eae85c18c65099b3c39a6b0a878c215441034
[11]:
  https://github.com/da-moon/podinfo/commit/9f50cc4442c91f5e38f07fdcc4258f3b8858c4a9
[12]:
  https://github.com/da-moon/podinfo/commit/f61c8359d544844e338df23b7474dbfa7c3a7106
[13]:
  https://github.com/da-moon/podinfo/commit/6999c69267506187c46538de68cfc66e150619e7
[14]:
  https://github.com/da-moon/podinfo/commit/3ad11d9c7bd206f9c46801d4f4c7f7b697d087b4
[15]:
  https://github.com/da-moon/podinfo/commit/d3178beb46cf01144cd02a6066d0258717e16ca7
[16]:
  https://github.com/da-moon/podinfo/commit/5f696a163155826d920ee4d0e6633512b226b630
[17]: https://github.com/da-moon/podinfo/compare/v0.8.1...v0.8.2
[18]:
  https://github.com/da-moon/podinfo/commit/868a12de21afde8e60164986f80f8f8521f3c4bd
[19]:
  https://github.com/da-moon/podinfo/commit/2e7d3fa4fb5668fb0467daf21ce303d6cc03852d
[20]:
  https://github.com/da-moon/podinfo/commit/06a0f867d367d1b19edf286d8c0c3248fea81bfa
[21]:
  https://github.com/da-moon/podinfo/commit/e8be68d1ebbdd270bd38ae16eaf009110eb9a124
[22]:
  https://github.com/da-moon/podinfo/commit/6f1f2a85003e07cc3d08de2f50671d2f893ded5e
[23]: https://github.com/da-moon/podinfo/compare/v0.8.0...v0.8.1
[24]:
  https://github.com/da-moon/podinfo/commit/952275560362291b3aca46793562f9b76196bf90
[25]:
  https://github.com/da-moon/podinfo/commit/7ec17481025b66ff69c94d4904f11ff2ddb3d544
[26]: https://github.com/da-moon/podinfo/compare/v0.7.0...v0.8.0
[27]:
  https://github.com/da-moon/podinfo/commit/344ff9ae23eee48c992fdc4e663ea0fedd324708
[28]:
  https://github.com/da-moon/podinfo/commit/68e97799e8952aa04bd0d692943c494377c314c6
[29]:
  https://github.com/da-moon/podinfo/commit/a3b0242d50158bd386735581ab8e1dd22cc65b20
[30]:
  https://github.com/da-moon/podinfo/commit/efb738cc1c46fd084a577426d275ff65337a145c
[31]:
  https://github.com/da-moon/podinfo/commit/13a50326a6671a4246854d5d948182aed907fff2
[32]: https://github.com/da-moon/podinfo/compare/v0.6.3...v0.7.0
[33]:
  https://github.com/da-moon/podinfo/commit/8e57115cc8778e3ece6885587881b4c6115d3307
[34]:
  https://github.com/da-moon/podinfo/commit/a5ab3a891ccfde2c4738da6ddcc79bee4805c182
[35]:
  https://github.com/da-moon/podinfo/commit/d12f7db17fdd6722aa6eec487928ac441c31ba11
[36]:
  https://github.com/da-moon/podinfo/commit/07c94d51c9543630f9436f238585b70429dcdc60
[37]:
  https://github.com/da-moon/podinfo/commit/3e0e1efd2247065243b8f5c9fde5c0416d52f311
[38]:
  https://github.com/da-moon/podinfo/commit/d0e25cbfb7b7ec66ce6b8015755bed6ad000b1b4
[39]:
  https://github.com/da-moon/podinfo/commit/8b5e5d2034c96183020879cbe910e86a68909c21
[40]: https://github.com/da-moon/podinfo/compare/v0.6.2...v0.6.3
[41]:
  https://github.com/da-moon/podinfo/commit/ab703346321ed38738d733022f5ad32a3ae256f6
[42]:
  https://github.com/da-moon/podinfo/commit/9935643552c94e1ddf9614315caaac0f27321045
[43]:
  https://github.com/da-moon/podinfo/commit/f154df8c72d4e3fa5da7ad46e5bb8155b4cc793a
[44]:
  https://github.com/da-moon/podinfo/commit/9de9d35a1bf70526e079cc701a46a8aebf51ec52
[45]:
  https://github.com/da-moon/podinfo/commit/61739b5b77ebcb89ee5e8a7fcbc7232de0bddee0
[46]: https://github.com/da-moon/podinfo/compare/v0.6.1...v0.6.2
[47]:
  https://github.com/da-moon/podinfo/commit/b7cb2bac6158e70794aeac4f3c70821b13fa1441
[48]:
  https://github.com/da-moon/podinfo/commit/6696e262bf7cf0790f3053c7d04b9dd74a9fe91b
[49]:
  https://github.com/da-moon/podinfo/commit/15a177c112e32b2dce6bf91556d145f2ae93109b
[50]:
  https://github.com/da-moon/podinfo/commit/26257c2f09b4272339f28a7d507b5354d3636df9
[51]:
  https://github.com/da-moon/podinfo/commit/1b6e7f10ffa32cbb5db368f1d8830ee638c6d634
[52]: https://github.com/da-moon/podinfo/compare/v0.6.0...v0.6.1
[53]:
  https://github.com/da-moon/podinfo/commit/f2f2164ae1858a9f7c356e1db8492b10ad2fc780
[54]:
  https://github.com/da-moon/podinfo/commit/fb15f041bc06c7ef1d7125af2df4df456ad6449a
[55]:
  https://github.com/da-moon/podinfo/commit/86f9e9d84c5a0aaf787ea52425631997ac98f1f0
[56]:
  https://github.com/da-moon/podinfo/commit/1e8dc7edd45ada829f64c8a771a216fe6b0f75cb
[57]:
  https://github.com/da-moon/podinfo/commit/f878c09e3d386081c7d695280b93cb8418445535
[58]:
  https://github.com/da-moon/podinfo/commit/cf372a5d5776b62452c58a9b3061e3f49f9a83e5
[59]: https://github.com/da-moon/podinfo/compare/v0.5.0...v0.6.0
[60]:
  https://github.com/da-moon/podinfo/commit/489db81e31deb74177750804a5e1a524194e6f26
[61]:
  https://github.com/da-moon/podinfo/commit/47b687dbc892100cfde86317431cd064b538bb89
[62]:
  https://github.com/da-moon/podinfo/commit/76a07eb02a9613f0a844d4dca2080d5086703765
[63]:
  https://github.com/da-moon/podinfo/commit/cc54fcfd488dc1fda7dffe332f676a89616cfcaa
[64]:
  https://github.com/da-moon/podinfo/commit/d26c4060d4dd22c3eb332e845bf217260788d796
[65]:
  https://github.com/da-moon/podinfo/commit/27323064403f94604111d257e64a450facb42867
[66]:
  https://github.com/da-moon/podinfo/commit/c313777862a30b59c4e69a58c62a20b9a166700d
[67]:
  https://github.com/da-moon/podinfo/commit/893a1e28050cb5a76850868cd46915159845c89f
[68]:
  https://github.com/da-moon/podinfo/commit/fc63a468f51d8ed52a8423a48ca00059d4cc4e83
[69]:
  https://github.com/da-moon/podinfo/commit/9e7f04ad39bea85318954a33d5fe6e6660cab628
[70]:
  https://github.com/da-moon/podinfo/commit/bff048aa3970ec87ae208a18a2a9bcb96f91fe8f
[71]: https://github.com/da-moon/podinfo/compare/v0.4.1...v0.5.0
[72]: https://github.com/da-moon/podinfo/compare/v0.4.0...v0.4.1
[73]:
  https://github.com/da-moon/podinfo/commit/5a9dc8e484995ee7b0f22e908c4e22bb31d1a95d
[74]:
  https://github.com/da-moon/podinfo/commit/f39830b2ab970d7528d40d0633ca61ff4eb2ad6b
[75]:
  https://github.com/da-moon/podinfo/commit/bf469a360fa978c971e00dbd3c9b5ce7b0d5fd87
[76]:
  https://github.com/da-moon/podinfo/commit/d22dc784cf87bfb439ac10040bf11aec56de408c
[77]:
  https://github.com/da-moon/podinfo/commit/e260245191ed449721c2fcf383e1d0ab05c240f7
[78]:
  https://github.com/da-moon/podinfo/commit/d7d5af78f46b4c34cd2b29611e5c3e360921cc37
[79]:
  https://github.com/da-moon/podinfo/commit/9d0b55f4629a3308f8d4e47e8154b081fb662029
[80]: https://github.com/da-moon/podinfo/compare/v0.3.0...v0.4.0
[81]:
  https://github.com/da-moon/podinfo/commit/10cef5d06baf2ad5668958de72396d7e1db187f7
[82]:
  https://github.com/da-moon/podinfo/commit/989f94806815ad67a097c60d49445a2b4642600e
[83]:
  https://github.com/da-moon/podinfo/commit/9871e64b8d6c9443f942c9dc70a5e139e0b033ed
[84]:
  https://github.com/da-moon/podinfo/commit/019355e452139b0b5b1cd7bf91feada7f12c70e2
[85]:
  https://github.com/da-moon/podinfo/commit/489218a1f2765a0b1d0a797fd7bc580f89f9a155
[86]:
  https://github.com/da-moon/podinfo/commit/239a061f4559e772d17adcda711179b6e5c3e88a
[87]: https://github.com/da-moon/podinfo/compare/v0.2.2...v0.3.0
[88]:
  https://github.com/da-moon/podinfo/commit/2d98b6a285beb7af80d3cd7e931cfc57a052e5c3
[89]:
  https://github.com/da-moon/podinfo/commit/958bc4765b272cd9229895381fa7248b7cf02e22
[90]:
  https://github.com/da-moon/podinfo/commit/0ef5a8f124d2833c63b109341a83813b905250f4
[91]:
  https://github.com/da-moon/podinfo/commit/877b7e49e5e675243323a458324fa0d103661ef4
[92]:
  https://github.com/da-moon/podinfo/commit/edc489e189bc852ad8afb6f8e2a419c4d5c9387d
[93]:
  https://github.com/da-moon/podinfo/commit/b66f28b69e76ad9a2dce8a90d29c26694fc90ee3
[94]: https://github.com/da-moon/podinfo/compare/v0.2.1...v0.2.2
[95]:
  https://github.com/da-moon/podinfo/commit/dead392b5dfd771e7f22d21a1d430e593ab87bb1
[96]:
  https://github.com/da-moon/podinfo/commit/b56a4009c45a0cbdbfae9bf37a6ebfb2985cc60c
[97]:
  https://github.com/da-moon/podinfo/commit/9356de010e16f533115cb79cc81c96ec46659b3a
[98]:
  https://github.com/da-moon/podinfo/commit/faeaf50147da8b4f20c72c9540903ffb2200e9f6
[99]:
  https://github.com/da-moon/podinfo/commit/65c6272690daa2c737c7029ef83a7e51664ce3d4
[100]: https://github.com/da-moon/podinfo/compare/v0.2.0...v0.2.1
[101]:
  https://github.com/da-moon/podinfo/commit/ae99b84a0a0e50044cfdd1e0c4ca23a2c9d06529
[102]:
  https://github.com/da-moon/podinfo/commit/a3851e663646f4f8ea5b7fdf25d98cca57eac868
[103]:
  https://github.com/da-moon/podinfo/commit/62974921e7bcc74d979d8044a21ceea768ed1da6
[104]:
  https://github.com/da-moon/podinfo/commit/3e7a28f9b4cf713cb7be3adfb4fd13b5f9d46da7
[105]:
  https://github.com/da-moon/podinfo/commit/fcdb30be2d06d616635d708a93eb8858b078b13c
[106]:
  https://github.com/da-moon/podinfo/commit/ce5be1d97bea3fc2641aa42857489b4df8860993
[107]:
  https://github.com/da-moon/podinfo/commit/b39a633efa7679a60560f96d835e10e074714fe5
[108]:
  https://github.com/da-moon/podinfo/commit/32376d6c5212da3f2f039e02aae3d9ac4524bde0
[109]:
  https://github.com/da-moon/podinfo/commit/ac1c0ee09913d8cc15ffc7da95381202519c07f8
[110]:
  https://github.com/da-moon/podinfo/commit/5e8ae1e07e0756e68fd40b141b5fb003858c5ec6
[111]:
  https://github.com/da-moon/podinfo/commit/4753ee2680b6770be50927164b2da4634482beb6
[112]:
  https://github.com/da-moon/podinfo/commit/2ec37fc20c165892cf2c9c68f2af68912eacc5b2
[113]:
  https://github.com/da-moon/podinfo/commit/f0efeb584a6014d973171dbc69f8c52cefd53c34
[114]: https://github.com/da-moon/podinfo/compare/v0.1.1...v0.2.0
[115]:
  https://github.com/da-moon/podinfo/commit/229f6ef6ee3ba17e7ad9d98e48fd30f6651484f0
[116]:
  https://github.com/da-moon/podinfo/commit/316c88d498470a1f1713b03e55e72acf8efa2f47
[117]:
  https://github.com/da-moon/podinfo/commit/fa4e4e0ca8e456eb05236a0a34ed3d888b22634a
[118]:
  https://github.com/da-moon/podinfo/commit/6975cebbfa43b6302b966b25a14d39178dde0e33
[119]:
  https://github.com/da-moon/podinfo/commit/0ae4c75ed64faf7e75d01dd6bed7592ba6b8d78c
[120]:
  https://github.com/da-moon/podinfo/commit/1fb4903412a18b8c72be1f1ed8f26302095479e5
[121]:
  https://github.com/da-moon/podinfo/commit/2af2f2519477a0f918713ebb84661a13fbc7320d
[122]:
  https://github.com/da-moon/podinfo/commit/a34f22d88cf9cbc2affdfdfb2a5fe92bf5a76d72
[123]:
  https://github.com/da-moon/podinfo/commit/4a735581212808bd7474cff0d971c8a8f08b1cd8
[124]:
  https://github.com/da-moon/podinfo/commit/8ea885513829cb69b379f4ed504921ea585e8957
[125]:
  https://github.com/da-moon/podinfo/commit/f81e20564262363a4bcfedee4664d93554b840a3
[126]:
  https://github.com/da-moon/podinfo/commit/c45a277c0c42c0101a94e59e8be27e0fb0b4dbd1
[127]:
  https://github.com/da-moon/podinfo/commit/c62180d42b8692b96e3933e9a59873e635455f08
[128]:
  https://github.com/da-moon/podinfo/commit/a989ce14ea993552bec8654521ece8b2307a837f
[129]:
  https://github.com/da-moon/podinfo/commit/5b226c3fbc0c939ce2ed28dbff92d42ae43be0f6
[130]:
  https://github.com/da-moon/podinfo/commit/7f85163ffaf583b762337b4123f5f833d11b87fe
[131]:
  https://github.com/da-moon/podinfo/commit/db54578cd8a6d8e35ca2c5c32f79a804fbe1233e
[132]:
  https://github.com/da-moon/podinfo/commit/2ac7d672cf635857b3e0d37d9de254aee3e42d12
[133]:
  https://github.com/da-moon/podinfo/commit/14fcb5748e6e54a6d1de4c2268f722900368424f
[134]:
  https://github.com/da-moon/podinfo/commit/7c7ecff03b89830c5f953444c0212ee24be31d7f
[135]:
  https://github.com/da-moon/podinfo/commit/777320ce11817c1e51cdb31bfa0769b29f0ceda4
[136]:
  https://github.com/da-moon/podinfo/commit/9788dfb8d3f123a55fc6f749674fd632d531de22
[137]:
  https://github.com/da-moon/podinfo/commit/92c17f302afaa44b2fd5eada40cfbe4360ec7241
[138]:
  https://github.com/da-moon/podinfo/commit/c9c5f7dfc082a281379bd0841c15b1de493bbb66
[139]:
  https://github.com/da-moon/podinfo/commit/c46195b2afc428cf3956fe3a474648b874ecd405
[140]:
  https://github.com/da-moon/podinfo/commit/d270881411ec57bf9f33286531ae4a7a9a18fdad
[141]:
  https://github.com/da-moon/podinfo/commit/66fd066549131258ce447316ba728a3009d5b83f
[142]:
  https://github.com/da-moon/podinfo/commit/6efbf25b42a1681ea7234d4bd6381aee0a40c469
[143]:
  https://github.com/da-moon/podinfo/commit/14a6fe0aa73f4d913156005dcbed34f70c9e2540
[144]:
  https://github.com/da-moon/podinfo/commit/ad05fc911ecd35b5c3e6a2238d5ab458acbfe8cf
[145]: https://github.com/da-moon/podinfo/compare/v0.1.0...v0.1.1
[146]:
  https://github.com/da-moon/podinfo/commit/e1bdb33b8719f89952e72518d5719952595db69a
[147]:
  https://github.com/da-moon/podinfo/commit/69d5e05a3d9f059fa608adf9d7d0915f96eb5a2b
[148]: https://github.com/da-moon/podinfo/compare/v0.0.1...v0.1.0
[149]:
  https://github.com/da-moon/podinfo/commit/ac8c96ccdaae2821bb002f18a705970d242d2dbf
[150]:
  https://github.com/da-moon/podinfo/commit/f264be1b6513c0e77b9667a02380cc4ed31594ef
[151]:
  https://github.com/da-moon/podinfo/commit/ae48e9d0e165c5b3ea3d774252f17860d9880887
[152]:
  https://github.com/da-moon/podinfo/commit/1c416f514caf64d5fa4e8eddd1e38b50bebe59dd
[153]:
  https://github.com/da-moon/podinfo/commit/73548a0f01a6dd4791e797866e9db587f317906d
[154]:
  https://github.com/da-moon/podinfo/commit/016f07c3bdeeb8ff5bb0eb5eefd5a016c16ebfd3
[155]:
  https://github.com/da-moon/podinfo/commit/133a5a8d61214cd437ab47431b7635e034462c4b
[156]:
  https://github.com/da-moon/podinfo/commit/a0b8f714a9f00047ba47e03185cc6eda821bc26f
[157]:
  https://github.com/da-moon/podinfo/commit/9bc13411d8edb4c995d0accfe936b597ad7ad74d
[158]:
  https://github.com/da-moon/podinfo/commit/a9eed4e11ac5b99dfc6621804d12c281c0fa6034
[159]:
  https://github.com/da-moon/podinfo/commit/11b80556c2492b622c6559eb555b6440fb2568c3
[160]:
  https://github.com/da-moon/podinfo/commit/fc5852ce35309c406992d2e79eeca1c8405f4c24
[161]:
  https://github.com/da-moon/podinfo/commit/7e25e0a5c130c92d17b33b9ada378c5d781b9219
[162]:
  https://github.com/da-moon/podinfo/commit/63d16c41732de0c05456d47a7861f330db1bf8e5
[163]:
  https://github.com/da-moon/podinfo/commit/5a0bb540b580c4774163c46ee5cc852348ebb39e
[164]:
  https://github.com/da-moon/podinfo/commit/d35f51366f156c455e5139ef3a947b49f3a028fb
[165]:
  https://github.com/da-moon/podinfo/commit/a8993b72b655b79ad9b978bc263d619d6066c342
[166]:
  https://github.com/da-moon/podinfo/commit/add524034530d1f143a5a0e775fe9a8db7cafedb
[167]:
  https://github.com/da-moon/podinfo/commit/06f7e413dd4721a45065a1f15b20f18abb99f9a2
[168]:
  https://github.com/da-moon/podinfo/commit/0e5f809c7c4739d5171dc5b816a91f27f34c2dd7
[169]:
  https://github.com/da-moon/podinfo/commit/2ac0c5e7abcde72f3e29335ce17c2311b96cbcf7
[170]:
  https://github.com/da-moon/podinfo/commit/efe22ac1a919947a20abfd752f9aadbd1bb5561a
[171]:
  https://github.com/da-moon/podinfo/commit/533f465a1b260a877450a8353d52e7fea444811c
[172]:
  https://github.com/da-moon/podinfo/commit/985c2cf10bac9c5eb3a7c6b02d4a713abdc0734c
[173]:
  https://github.com/da-moon/podinfo/commit/02ce2eb9d4c4936aca40fb2cda3deebc73aad47e
[174]:
  https://github.com/da-moon/podinfo/commit/f52d60804993c8265fce153174471fe1e95d095f
[175]:
  https://github.com/da-moon/podinfo/commit/b9311001df80123f8d93b933cf68d1b25df132ce
[176]:
  https://github.com/da-moon/podinfo/commit/006af3c587eebaefe9fe3d9849220e38750ca349
[177]:
  https://github.com/da-moon/podinfo/commit/6d9e725e39c685529b98478d662304e069ca37cf
[178]:
  https://github.com/da-moon/podinfo/commit/9db2bb8f3175a5e99c91853904a1bb78d97e3f85
[179]:
  https://github.com/da-moon/podinfo/commit/747b4f4ad9c572e2139d0bf9809ad2e9e6b5b286
[180]:
  https://github.com/da-moon/podinfo/commit/c397414d9ce68237d977906a041bdb2a17340492
[181]:
  https://github.com/da-moon/podinfo/commit/7a6a5d91194ed7bfeaf6e9b180b940c0190d2428
[182]:
  https://github.com/da-moon/podinfo/commit/495f420a04864c48f43f75092dd9b1dd2df9fbb9
[183]:
  https://github.com/da-moon/podinfo/commit/9a5964e34021606c24d5128050a1364096a87903
[184]:
  https://github.com/da-moon/podinfo/commit/5ca9bde78929d67a800c0b01225dc0b50de9537f
[185]:
  https://github.com/da-moon/podinfo/commit/74090a2cd3fb42b78b4f568a64c09398ca2f694f
[186]:
  https://github.com/da-moon/podinfo/commit/0c641bed2cd1666f2873bb0ce39df75960b8ddb6
[187]:
  https://github.com/da-moon/podinfo/commit/385d829da17c0b9d5a51e5add76cfe0df3db7eca
[188]:
  https://github.com/da-moon/podinfo/commit/2e452248570e3ae6b8fe25b2525e5c8e0a15fb6b
[189]:
  https://github.com/da-moon/podinfo/commit/55e6ae7144c137be0ac3383431908d8fd241b0a7
[190]:
  https://github.com/da-moon/podinfo/commit/2011e4f64d01c9641973d9e7aab54d27eb9e70f3
[191]:
  https://github.com/da-moon/podinfo/commit/ea4cc927e6b2ec5e2244e7e3d769b645bdf9db98
[192]:
  https://github.com/da-moon/podinfo/commit/168e63fb281dd963d4fedc3b0234af9de0511366
[193]:
  https://github.com/da-moon/podinfo/commit/e3e3d0de719b84dc0ed1751a1524c2b1c87723fe
[194]:
  https://github.com/da-moon/podinfo/commit/d2ce4f17f4d20ac16767b483a657acceb0fe8ad8
[195]:
  https://github.com/da-moon/podinfo/commit/79f5f6d94403317ee0040958ad02ff177be8985f
[196]:
  https://github.com/da-moon/podinfo/commit/f9bc049b3e3436706bde11b5d71f05612c8ee2ff
[197]:
  https://github.com/da-moon/podinfo/commit/9ae11b2658110e4695059a17fc27f17b0ea4c324
[198]:
  https://github.com/da-moon/podinfo/commit/ee99516683eea710d9a3503d35960670e163566b
[199]:
  https://github.com/da-moon/podinfo/commit/927e64b33134b86005ecd859c202c03c1ef38600
[200]:
  https://github.com/da-moon/podinfo/commit/f494de7904035ae5a2a412b853d8d67409a3a7e2
[201]:
  https://github.com/da-moon/podinfo/commit/8c3577e3792424c41befc47696ef2b07c0c2f7aa
[202]:
  https://github.com/da-moon/podinfo/commit/3974f9545008b57bd0cc276ec1b0ee6640905cb1
[203]:
  https://github.com/da-moon/podinfo/commit/891fa92152f41ebcca17b8efa48daff2355bbe53
[204]:
  https://github.com/da-moon/podinfo/commit/dc9c0cc25a0f009ea51c3736248611efff930210
[205]:
  https://github.com/da-moon/podinfo/commit/a9354fe0f7deb386b8a0f62f2a0c23ea5b099477
[206]:
  https://github.com/da-moon/podinfo/commit/03877c675e283e9a00db1b6c09b5fc3486ebf949
[207]:
  https://github.com/da-moon/podinfo/commit/c093b2a79d3e495eb6a5d170d3100f8c481c8a3d
[208]:
  https://github.com/da-moon/podinfo/commit/a3c1ebf16a4f5e0c92bb98a483bc4f09d9ddc8be
[209]:
  https://github.com/da-moon/podinfo/commit/34a459788c8b449d7957cb11fd3223fbab4168be
[210]:
  https://github.com/da-moon/podinfo/commit/4ac385ca45dee22b5c208289e584c6dc77e68fc1
[211]:
  https://github.com/da-moon/podinfo/commit/f1264aa5f3f1e01681998249824adcdc714a5f24
[212]:
  https://github.com/da-moon/podinfo/commit/09e59d1ec1c1c96dd41d5891eab2e32424129be8
[213]:
  https://github.com/da-moon/podinfo/commit/e066d10c03acf681ae25f017bb76817b31c18fe0
[214]:
  https://github.com/da-moon/podinfo/commit/a39b7c2da11cc4f6b9a6c88579725e4d4d376bee
[215]:
  https://github.com/da-moon/podinfo/commit/489996ce7a195f28c9175c8ab3889524b25bce23
[216]:
  https://github.com/da-moon/podinfo/commit/a29c5d7b0674babb6273f852568b38ee7e265857
[217]:
  https://github.com/da-moon/podinfo/commit/2bc4ce137a568e5f2670ffa1818d185a0e808871
[218]:
  https://github.com/da-moon/podinfo/commit/1a67ad670c4e6367440f24440565e086b3dc1a43
[219]:
  https://github.com/da-moon/podinfo/commit/76eb234e4aa085d6309ebaf74fbf01d329dc74ab
[220]:
  https://github.com/da-moon/podinfo/commit/c188be8e647bffa90f7416d07721ce82bec1c819
[221]:
  https://github.com/da-moon/podinfo/commit/58ed4ff6d80d36426d878cff2df14bdbdf4e4fb9
[222]:
  https://github.com/da-moon/podinfo/commit/4a98993f0aebc04358446cc43c004c97849f5892
[223]:
  https://github.com/da-moon/podinfo/commit/f8d487e2c8c4308acf55847774d45e84b2dc550f
[224]:
  https://github.com/da-moon/podinfo/commit/f70b48d1b76c0bc330430e03918595fdbf832830
[225]:
  https://github.com/da-moon/podinfo/commit/437e93c9bb8882d84f2fb393e193c3c4d7f75a8f
[226]:
  https://github.com/da-moon/podinfo/commit/9d115adcb11be99c6132245d0cc0e5b53309d3c2
[227]:
  https://github.com/da-moon/podinfo/commit/5b380a78fa9a7198b515d35d5cca5a4c41b5eca7
[228]:
  https://github.com/da-moon/podinfo/commit/a2331417e8d83ab0b541147e4948fe49f7c0e00f
[229]:
  https://github.com/da-moon/podinfo/commit/aef98ca27e70468e1b2b8bf744324da8cee743b5
[230]:
  https://github.com/da-moon/podinfo/commit/8196569d3f4a3f8739dfd1afb8db9d70d4b1c6c0
[231]:
  https://github.com/da-moon/podinfo/commit/1e976081ee00e280042f1df74ffd92cdce7c3def
[232]:
  https://github.com/da-moon/podinfo/commit/5dbb2e99754caa0031a8dd92d60a3161ca19b6a7
[233]:
  https://github.com/da-moon/podinfo/commit/fade8c6cb9fa8d95d5ae2f1b0f88e11a8b7b5ba4
[234]:
  https://github.com/da-moon/podinfo/commit/69918d748b7f6c52e565c820a380f8c40addde84
[235]:
  https://github.com/da-moon/podinfo/commit/e5403e6aee54bf47bbb1119d8d7abb678a2ac64b
[236]:
  https://github.com/da-moon/podinfo/commit/fd5e5e2caa33dcbdf5891018fa558366ab5476d7
[237]:
  https://github.com/da-moon/podinfo/commit/979298e4bbc0c7ab287b872ffe9af696afd874be
[238]:
  https://github.com/da-moon/podinfo/commit/3ddbd38d57f2138c941de21c3595326c839f6e46
[239]:
  https://github.com/da-moon/podinfo/commit/76ab9c82deb31a4a0f8bf7893eedfb58fcce3132
[240]:
  https://github.com/da-moon/podinfo/commit/2a7f6c7c4309bf148d0d59a2dd909cc072716193
[241]:
  https://github.com/da-moon/podinfo/commit/829e8ef6a522b2cba4766e8a74430cf1287dbce9
[242]:
  https://github.com/da-moon/podinfo/commit/48e79407db3a8c7e7f5bca9e766c0a5474e4917b
[243]:
  https://github.com/da-moon/podinfo/commit/5eb4b1836c69ff2c2cfb8484440c55607794d9d9
[244]:
  https://github.com/da-moon/podinfo/commit/b1895bfa689cabedac7da6b68410222cd719f65c
[245]:
  https://github.com/da-moon/podinfo/commit/fa52696656338f964f4eeefd283d74d3e0161c90
[246]:
  https://github.com/da-moon/podinfo/commit/7d4928f426ee9ecbbe6559a28630428f6a9d7278
[247]:
  https://github.com/da-moon/podinfo/commit/8660a4d3daee904c81ead3f010bb6c198adda404
[248]:
  https://github.com/da-moon/podinfo/commit/540b954af013719627265bd31ddfa4fd873c6435
[249]:
  https://github.com/da-moon/podinfo/commit/df8d8d1a1bdd417ad8ca71a2f2951dbfcc57e06a
[250]:
  https://github.com/da-moon/podinfo/commit/c4cd6da3968f6611ff020bc6b78a7a19a6984fa2
[251]:
  https://github.com/da-moon/podinfo/commit/93e80f2991001187e5646ed01484af5aaf9ff3c4
[252]:
  https://github.com/da-moon/podinfo/commit/a9e2bfbb42cf97eef57ef4ecda5adde0d1346a40
[253]:
  https://github.com/da-moon/podinfo/commit/966607f0c7d71b311e0907ecc6a7eabdcd783dcd
[254]:
  https://github.com/da-moon/podinfo/commit/d4bb8196b2244439e0e462dad07df0422fa6a79b
[255]:
  https://github.com/da-moon/podinfo/commit/89c2bae73efdae1c55632e6824e1999b321ea010
[256]:
  https://github.com/da-moon/podinfo/commit/7089f0bf892022d41989eec4a13a90d2c09badc6
[257]:
  https://github.com/da-moon/podinfo/commit/25b84796290bdf2118e4b562e57808ebee17ec83
[258]:
  https://github.com/da-moon/podinfo/commit/1c62db812dc7c32352aaf7115e1c5be6caf17853
[259]:
  https://github.com/da-moon/podinfo/commit/e37bcd05eb9d5cdd715d70b3c26e1cfb6a9c473c
[260]:
  https://github.com/da-moon/podinfo/commit/50231ce1e3900bb5e32abacc3ec02fd89855da36
[261]:
  https://github.com/da-moon/podinfo/commit/f240c28e3aa9e3e75deda95eaa986ebc19984a5c
[262]:
  https://github.com/da-moon/podinfo/commit/5ac43ba89ca7144dad563a38c4eb93bd7a23b7e1
[263]:
  https://github.com/da-moon/podinfo/commit/92ec6b40f7350b5d9c37ed1df9b47305e3b80c4f
[264]:
  https://github.com/da-moon/podinfo/commit/2721074be8221e3d305f14fde36238124284ade5
[265]:
  https://github.com/da-moon/podinfo/commit/78a79654c2374fd336b2ed293b159801427a097d
[266]:
  https://github.com/da-moon/podinfo/commit/8c5778c37a4f6b44e0b55e86a42d0851e9127dfa
[267]:
  https://github.com/da-moon/podinfo/commit/6a417bd6d8d0ff801b23d9897d6c0121124db469
[268]:
  https://github.com/da-moon/podinfo/commit/2e74c8d513ac30ed7ffe748fa1cc40a97e2733c6
[269]:
  https://github.com/da-moon/podinfo/commit/b59a5d766d3fd36cabf94add3118911312eb4f11
[270]:
  https://github.com/da-moon/podinfo/commit/08b1ac60aea5b9ace3ac52d5b4dc4c89a28c674b
[271]:
  https://github.com/da-moon/podinfo/commit/ef340bbc8d27f3339dfb19b13f46adb82b0a4406
[272]:
  https://github.com/da-moon/podinfo/commit/4026b9e2732c596170c6cedfdad046cd230d8285
[273]:
  https://github.com/da-moon/podinfo/commit/deb5e1d039766ac02efa9819213c80ac81920ff0
[274]:
  https://github.com/da-moon/podinfo/commit/aa5683bad24ca093e47477b6250a51a6e86aadf8
[275]:
  https://github.com/da-moon/podinfo/commit/74afa35fdce53ee030589b261f06cf32efcc1c7c
[276]:
  https://github.com/da-moon/podinfo/commit/2bde4d84177b8a752185e7343fc93a55f02c2af4
[277]:
  https://github.com/da-moon/podinfo/commit/03bb13d60e877cdd543e2164828c96d606607d33
[278]:
  https://github.com/da-moon/podinfo/commit/a5db8af3fcb33cb00928a08b5298698ee58e7450
[279]:
  https://github.com/da-moon/podinfo/commit/d02b807abc37384811ecbea6ff080aaa17cc21b5
[280]:
  https://github.com/da-moon/podinfo/commit/c367eaedeec6283bcc966094adabf9e6fe70f3e8
[281]:
  https://github.com/da-moon/podinfo/commit/6aee1b8a085cc4c0b3073c189470f2ea32914839
[282]:
  https://github.com/da-moon/podinfo/commit/f84a9ebfa68a15d779910ff2b3bedf191ef43072
[283]:
  https://github.com/da-moon/podinfo/commit/3781d547e78d0c0675fa217f3d6933ba9b734354
[284]:
  https://github.com/da-moon/podinfo/commit/b8d961de6f51feec0cadf482b070f2b1f4442725
[285]:
  https://github.com/da-moon/podinfo/commit/0d1588e2f07b3b02f8f00c8af71375b3a7862df0
[286]:
  https://github.com/da-moon/podinfo/commit/b223255795b644c6f8546041a254826439587412
[287]:
  https://github.com/da-moon/podinfo/commit/58aabd6c273c2139192fa9d1a893541930286840
[288]:
  https://github.com/da-moon/podinfo/commit/00df92c2f70710450781433c117a641a0df64f64
[289]:
  https://github.com/da-moon/podinfo/commit/9a78d0a16db5259902d07935f46e0ec8804a3858
[290]:
  https://github.com/da-moon/podinfo/commit/b817954a420f685dfd1db2db713d5d16a04ed245
[291]:
  https://github.com/da-moon/podinfo/commit/4b94227327f10a8bcca2e178325033a14ba0baf9
[292]:
  https://github.com/da-moon/podinfo/commit/f80992fbee627a72d99b0b6789d74d9088ebeb06
[293]:
  https://github.com/da-moon/podinfo/commit/9cc3843cd5ce4857dac264db43cbab1909672961
[294]:
  https://github.com/da-moon/podinfo/commit/7a17aad695a02969b7e206fc0e3ea73bef6ddf78
[295]:
  https://github.com/da-moon/podinfo/commit/40e593d79d0cc2ce56e4264369cf8915e7303ad1
[296]:
  https://github.com/da-moon/podinfo/commit/7e5375bba1f7a2a7cc54281d27f8e481560505e6
[297]:
  https://github.com/da-moon/podinfo/commit/960062dca5e4ccebb66762c0da8883b984c74165
[298]:
  https://github.com/da-moon/podinfo/commit/2a1ffd489a71459e35bbc6dbcf32111b09c127e1
[299]:
  https://github.com/da-moon/podinfo/commit/265da8cbbb7155de126bb584cd97f2063bd92135
[300]:
  https://github.com/da-moon/podinfo/commit/295aad0a4aa7beaac94389291514c06b3328f9ba
[301]:
  https://github.com/da-moon/podinfo/commit/9f0d0986e7abc63a8dba91cd0dbe69899ba48228
[302]:
  https://github.com/da-moon/podinfo/commit/9ef7b34f348e1244aaf5cb95ba366584821070af
[303]:
  https://github.com/da-moon/podinfo/commit/b0e08746f124e45f0836475b4f8bff12301ee56c
[304]:
  https://github.com/da-moon/podinfo/commit/d3233c89166be6061014360a8e0ba507a8743c01
[305]:
  https://github.com/da-moon/podinfo/commit/e801eadcb5e980758b9d2428a5d47990244e1d67
[306]:
  https://github.com/da-moon/podinfo/commit/a163f57dfa7edacfb33fcf92a0a81874f715945f
[307]:
  https://github.com/da-moon/podinfo/commit/d63c0a7a1a1d889a6665e67ab9c0e4ba73ba3ad4
[308]:
  https://github.com/da-moon/podinfo/commit/5d6533065761c4eaa7c232bd1132dcae79223d38
[309]:
  https://github.com/da-moon/podinfo/commit/41feade1f704a0647ab0b33422be1584f0bd410f
[310]:
  https://github.com/da-moon/podinfo/commit/87c51ba596feaf4dc341999362c41b2b5fc842c9
[311]:
  https://github.com/da-moon/podinfo/commit/8a083ac1bfb27994769adf207a3ad897172d92fd
[312]:
  https://github.com/da-moon/podinfo/commit/dffbbdced9a4bda445be2dedb486cbf6d713e354
[313]:
  https://github.com/da-moon/podinfo/commit/3895ee7e995db28e18e9c646f538c2ff5b855deb
[314]:
  https://github.com/da-moon/podinfo/commit/ab5c0b9ee149b1899d696b2d75903c9ee84d54e3
[315]:
  https://github.com/da-moon/podinfo/commit/f657f443d0cc2c02df40274a172efe1c50b3a107
[316]:
  https://github.com/da-moon/podinfo/commit/9706c76925065ed8bd02bb80ea84b755e66e1332
[317]:
  https://github.com/da-moon/podinfo/commit/108d9eb30d47b24f1594112143d509ae37ead56a
[318]:
  https://github.com/da-moon/podinfo/commit/3574a9ca02677d07332b8978408a43a12191a482
[319]:
  https://github.com/da-moon/podinfo/commit/a13ad40df0a4d334444c1cbda58ea0ab228f2acf
[320]:
  https://github.com/da-moon/podinfo/commit/0a9d0c1eca6fbfc92bee69069b537bf6709ef1c3
[321]:
  https://github.com/da-moon/podinfo/commit/83a078c35ca693ca2147d801512cb73c050491f9
[322]:
  https://github.com/da-moon/podinfo/commit/78abe82036aeb0a6b60c14edd335426fc2a9dbe3
[323]:
  https://github.com/da-moon/podinfo/commit/4e53b916c4fff92a2e00ac677d069238caf11b34
[324]:
  https://github.com/da-moon/podinfo/commit/23ab44b7176bc010e20c25374916ddd2159959a8
[325]:
  https://github.com/da-moon/podinfo/commit/2a7c7909e97f65161a70ffced480b3b9785bfa96
[326]:
  https://github.com/da-moon/podinfo/commit/c5057285bf86f27c2d4cb6695054fea0abf684be
[327]:
  https://github.com/da-moon/podinfo/commit/1090e11a1353a0d8756bcf677423ccb283e0135e
[328]:
  https://github.com/da-moon/podinfo/commit/197fad7f1f6f9dc12eafbe161f1f2cb0ed929529
[329]:
  https://github.com/da-moon/podinfo/commit/5cdbfab66287cc80a1b993f4f36dd93d108c8ea1
[330]:
  https://github.com/da-moon/podinfo/commit/6d23ed64ecf8713e014ef3feec0306f2bd899c7d
[331]:
  https://github.com/da-moon/podinfo/commit/cea7bee6e30457855359ad2bd6254f9049e4d0b1
[332]:
  https://github.com/da-moon/podinfo/commit/13e7f7daf1528b3ef49df1bbda37922fe786c026
[333]:
  https://github.com/da-moon/podinfo/commit/4b31bb451e2fe001ab6e3974cc91c078949c5c59
[334]:
  https://github.com/da-moon/podinfo/commit/49a1e335e29e8bc5acd59610fbc0c176d0f6bf6f
[335]:
  https://github.com/da-moon/podinfo/commit/70446b16d0c11691fc882f237a3a03c6d0aea447
[336]:
  https://github.com/da-moon/podinfo/commit/1f906c53039a763fc092c2cb80d2a66d5c2678aa
[337]:
  https://github.com/da-moon/podinfo/commit/13983aa943ce42e056d213691327b02c5c62335f
[338]:
  https://github.com/da-moon/podinfo/commit/942978f2e80199ec6b0085d4aafe622ef2876280
[339]:
  https://github.com/da-moon/podinfo/commit/26ea212aaa7e661d5b9f7d5a714ef000d88dde1e
[340]:
  https://github.com/da-moon/podinfo/commit/5344b78d92dc64f14b5f8cee99d91c205b8d9db0
[341]:
  https://github.com/da-moon/podinfo/commit/e55becaafbef76ba15d4c216b8c90678951a6298
[342]:
  https://github.com/da-moon/podinfo/commit/baf227034b6bad8a0dfc186fa1323ccb09f51660
[343]:
  https://github.com/da-moon/podinfo/commit/abe018f7335b1fb9f7dcd77f141e516c9e1722b7
[344]:
  https://github.com/da-moon/podinfo/commit/790990049584461133de87a4b42d96d855c376ed
[345]:
  https://github.com/da-moon/podinfo/commit/750e36dc16771beb96834a1728d0c3c287aa5892
[346]:
  https://github.com/da-moon/podinfo/commit/4c68a04bb689b97b028788f5946ef2300b2a1143
[347]:
  https://github.com/da-moon/podinfo/commit/83e55344dc1b41cb15fe893ecdd8820a936131be
[348]:
  https://github.com/da-moon/podinfo/commit/e238d0e382e348eea8068609332eaac20aa7af25
[349]:
  https://github.com/da-moon/podinfo/commit/dba5a78132f74695cbf453b51904fc89cf1b9830
[350]:
  https://github.com/da-moon/podinfo/commit/4124f1286c94978ad6946802c6272282e0dd6ef4
[351]:
  https://github.com/da-moon/podinfo/commit/c7d8513c870ba0386fddfc98da9e1ecb615d708b
[352]:
  https://github.com/da-moon/podinfo/commit/d60a2a973b219b9de5957ed0644154ee50637275
[353]:
  https://github.com/da-moon/podinfo/commit/eabe9eaec00dad2c25a78ae9eb18b91b5158b7ec
[354]:
  https://github.com/da-moon/podinfo/commit/0598dda6296422263c4dfa2d62dae27a2e9a0b29
[355]:
  https://github.com/da-moon/podinfo/commit/ffb9e6bbbd3ae65e4b4d6c12d6eec8469543a5f0
[356]:
  https://github.com/da-moon/podinfo/commit/1d6b435537e5c92f971bc10a7d41698c3cd3c296
[357]:
  https://github.com/da-moon/podinfo/commit/b3ebabc96f03da5b32d62337d740dd5165076b29
[358]:
  https://github.com/da-moon/podinfo/commit/c3255c27f705e32888700d81d6655f34a21d8ee3
[359]:
  https://github.com/da-moon/podinfo/commit/588c83a94df3d7f773d1e1e078ab7df3a7b7ea6a
[360]:
  https://github.com/da-moon/podinfo/commit/1b3cfbe3be38c7079cbd7b8e96494e53ee67cfba
[361]:
  https://github.com/da-moon/podinfo/commit/4aeaebb2cb4457c4b12a9152312009668ee4356f
[362]:
  https://github.com/da-moon/podinfo/commit/01cd7be2e05f1a6e5bebeb8ffbe7bd80eaac1dc8
[363]:
  https://github.com/da-moon/podinfo/commit/3382f775bcb1a4635c55b9ac3de07eaaccd94bf5
[364]:
  https://github.com/da-moon/podinfo/commit/e262bfad3275b9929c001b87838bcbd9afc4b005
[365]:
  https://github.com/da-moon/podinfo/commit/40a06378911255c22a887c329d811b6f5266652f
[366]:
  https://github.com/da-moon/podinfo/commit/d8670aaafd6217c021785fb5f56a60afe97a21c0
[367]:
  https://github.com/da-moon/podinfo/commit/8dc7e7c3da53d7e1acfa1192635bb4b61b8244be
[368]:
  https://github.com/da-moon/podinfo/commit/14209fdd6b0137261ca4bb8ae9196e0ce195cd02
[369]:
  https://github.com/da-moon/podinfo/commit/012682644ff38227dba27e97a8728fa71f611bd5
[370]:
  https://github.com/da-moon/podinfo/commit/39b68ceabe3b7e134e4d04a4418946bf6398c1cf
[371]:
  https://github.com/da-moon/podinfo/commit/47dc4e8bc33b1e3b9c15c4678514a2e312602ad9
[372]:
  https://github.com/da-moon/podinfo/commit/26a247ae64322896579f6aaf26ac6d281eea3932
[373]:
  https://github.com/da-moon/podinfo/commit/351a32a326a6f77de8c06b9bc4d6800912da4b78
[374]:
  https://github.com/da-moon/podinfo/commit/3ca58f4b0e46a911dc13d8af89901c0c98a81ec7
[375]:
  https://github.com/da-moon/podinfo/commit/852e53eabb1d35df4a05a85af480166d20e94f3e
[376]:
  https://github.com/da-moon/podinfo/commit/eb8efa2ee230c40490df576005724c743e866620
[377]:
  https://github.com/da-moon/podinfo/commit/e540c340817a8c70395109ef819e0b75132c204a
[378]:
  https://github.com/da-moon/podinfo/commit/2754846dcd2523dbeba9ee85af8233352208d438
[379]:
  https://github.com/da-moon/podinfo/commit/5cfd1f9dfaba7615a067e08264d902f0ffc87441
[380]:
  https://github.com/da-moon/podinfo/commit/05d97b0c6de71d7ea4626080c993975372d22b81
[381]:
  https://github.com/da-moon/podinfo/commit/2e5683972180b3ae8cddcea4c4b22b325bd66723
[382]:
  https://github.com/da-moon/podinfo/commit/221424224fcf93635a074a970db268560140c3d2
[383]:
  https://github.com/da-moon/podinfo/commit/184f86e3c08b5e65fb243059d7626ba8ae4d5450
[384]:
  https://github.com/da-moon/podinfo/commit/8669e4832453e623b068bd9ecbce1df44483834c
[385]:
  https://github.com/da-moon/podinfo/commit/d9813b6a392dfd7c2f6a02747928f2091c27ea20
[386]:
  https://github.com/da-moon/podinfo/commit/0ee99d80b19dc9d83d90c283aadf5f832c4835b1
[387]:
  https://github.com/da-moon/podinfo/commit/47c8d9d3d3e87b2e0c7474052460c285fad4df46
[388]:
  https://github.com/da-moon/podinfo/commit/4ca9c5c58b3de65774caaea1c8a26e3e451f4503
[389]:
  https://github.com/da-moon/podinfo/commit/fb9c78ae6233e95fa3d6f242054d6ad3ce61e391
[390]:
  https://github.com/da-moon/podinfo/commit/f574bf3e6fcd7a6c1d7f6a4e179f9b348d5852da
[391]:
  https://github.com/da-moon/podinfo/commit/124347227b3af638363feb824334f144db2b99bc
[392]:
  https://github.com/da-moon/podinfo/commit/e5c520677feec851e93687d6e0f7e56faa8801f5
[393]:
  https://github.com/da-moon/podinfo/commit/b053a644e2245ddcf2b3df9114667d29787820d8
[394]:
  https://github.com/da-moon/podinfo/commit/28764fbf86eb372be8320573d2e839db902dfa33
[395]:
  https://github.com/da-moon/podinfo/commit/33b956f82a6e2a8f2a3ad44c5955ac42f7eb4cf3
[396]:
  https://github.com/da-moon/podinfo/commit/81aab57c00158181b1e2cc8ae8659d850c8589de
[397]:
  https://github.com/da-moon/podinfo/commit/9b5d2d8ba4da857eac1a9bafc51beaa524a6a206
[398]:
  https://github.com/da-moon/podinfo/commit/d8c6c38b9e61b132b05735a5ddb149dc4536c251
[399]:
  https://github.com/da-moon/podinfo/commit/2a68b485cedd24efc2bd3a6654f89868d76ee209
[400]:
  https://github.com/da-moon/podinfo/commit/07c21bb4c0b6e07e3e775c33d5b74f19dc110436
[401]:
  https://github.com/da-moon/podinfo/commit/0cb70f375eed7aa02a11352fd80124da98c579e0
[402]:
  https://github.com/da-moon/podinfo/commit/9c70c26e2fdcf3ece18445e23d5ef68b888314b2
[403]:
  https://github.com/da-moon/podinfo/commit/7fb7d1486e857077bc3be7b96d4f58c572330106
[404]:
  https://github.com/da-moon/podinfo/commit/8ffe08383a50ebd06bfeb43a3dbb41c5bb66f2b2
[405]:
  https://github.com/da-moon/podinfo/commit/86033e01235abb2b1d9d1a2e4e53a41f5aa56738
[406]:
  https://github.com/da-moon/podinfo/commit/52e78862355be6d413eba2ade4ef85299d7f3251
[407]:
  https://github.com/da-moon/podinfo/commit/84ef4a79ef952d11afae6197e355fd044214937d
[408]:
  https://github.com/da-moon/podinfo/commit/a07a063dd390f118e1c195b6f23cd05fc7fb2066
[409]:
  https://github.com/da-moon/podinfo/commit/ae05d7a512b6b3c6756933da5c95e49df1785e22
[410]:
  https://github.com/da-moon/podinfo/commit/1906bdca71f82bd93bed8061aaa7c8df3c919abf
[411]:
  https://github.com/da-moon/podinfo/commit/778b6f762c00d061db12ababd5b335055d321a05
[412]:
  https://github.com/da-moon/podinfo/commit/dd58811ce31a676e7fde593c33b52f5ed772476a
[413]:
  https://github.com/da-moon/podinfo/commit/763dcb74aa0059c2f4c0172ad7635dd1ee7e91b1
[414]:
  https://github.com/da-moon/podinfo/commit/7feeae804ae29f175e577fd110ce490cc993771e
[415]:
  https://github.com/da-moon/podinfo/commit/c0493f8453c2ac1391518e61bd18e94cc9d19d87
[416]:
  https://github.com/da-moon/podinfo/commit/0cebd31bc14b0a5ea8957d2b500c2404aafcf72b
[417]:
  https://github.com/da-moon/podinfo/commit/a89e0be49244d709a7acaee7e7a4e8dc6a9c3988
[418]:
  https://github.com/da-moon/podinfo/commit/607964a1e0ca05ed5453c08793acec88d2c6f107
[419]:
  https://github.com/da-moon/podinfo/commit/2d320e6b971971dd23676230d7cb5814e14730c0
[420]:
  https://github.com/da-moon/podinfo/commit/25bea21c6d047c318346fba7ce4df09fa84b3754
[421]:
  https://github.com/da-moon/podinfo/commit/2e7af353c4ad21745100296561a99b330ab3a686
[422]:
  https://github.com/da-moon/podinfo/commit/84e27cab63949069aae794d54b27868ebe5a1796
[423]:
  https://github.com/da-moon/podinfo/commit/5471c3438f86f3ab43e678d508c62d847bea8dc9
[424]:
  https://github.com/da-moon/podinfo/commit/49087435e73f6069ce83f77e681623b74c9ce090
[425]:
  https://github.com/da-moon/podinfo/commit/fc1e8f45714b33d54b65d3fc8e20e54ae40cce44
[426]:
  https://github.com/da-moon/podinfo/commit/299d59e6eebd59ee80f58607126a9cf0b3ac35ca
[427]:
  https://github.com/da-moon/podinfo/commit/4e9b6769bde6cbafdb363102b6503082f1213b3f
[428]:
  https://github.com/da-moon/podinfo/commit/57b16f2c44ab5131566b6a9e98dec8cec21e82dd
[429]:
  https://github.com/da-moon/podinfo/commit/8c4b6c496f72c59005186f66b7cdf690facd2730
[430]:
  https://github.com/da-moon/podinfo/commit/b501fa4d46217f962c00025f2e4c303795d7c865
[431]:
  https://github.com/da-moon/podinfo/commit/755f5d2c95e41e651398561c19e0886f880e1394
[432]:
  https://github.com/da-moon/podinfo/commit/825e7c28d148ff7798292ac2a059814ed9b2b43e
[433]:
  https://github.com/da-moon/podinfo/commit/f55402bc7ce26d79b6df293b90ffe649b4fecaa9
[434]:
  https://github.com/da-moon/podinfo/commit/6ff23fe9d3dc9ea61bb2da5a59da6082a82293a4
[435]:
  https://github.com/da-moon/podinfo/commit/a160290b3961d9cea6b5b051c07553a15710f721
[436]:
  https://github.com/da-moon/podinfo/commit/fbfd5b767c26e4915dfedfe7c6b74add40f2b1ff
[437]:
  https://github.com/da-moon/podinfo/commit/502e7a515260b0e819f9ba331fcca7d31373c2b4
[438]:
  https://github.com/da-moon/podinfo/commit/7a098f47bebabb61ac0c2c533f5b2e59c6ee68c1
[439]:
  https://github.com/da-moon/podinfo/commit/b61d49c242404f9669e85914a7f43be646aeda85
[440]:
  https://github.com/da-moon/podinfo/commit/a8db8f73119d5d5781a1010dee5d47b28028eaa2
[441]:
  https://github.com/da-moon/podinfo/commit/713d5ab3ee719e5de919dc480baecd97e0e2d3cc
[442]:
  https://github.com/da-moon/podinfo/commit/2fbb351f0a99ad715eed6459d49d0f316a100d61
[443]:
  https://github.com/da-moon/podinfo/commit/7d5618d896e2f57d6f342f542711a363a7e24855
[444]:
  https://github.com/da-moon/podinfo/commit/8d9e1cd838e30daded0447817b928c7ab9f22712
[445]:
  https://github.com/da-moon/podinfo/commit/3ec3faf85068a45de8de8072497ca9a827744507
[446]:
  https://github.com/da-moon/podinfo/commit/7ab2d30bc7c26d0250ba91e665a6a543aa748065
[447]:
  https://github.com/da-moon/podinfo/commit/2f05faf322da5e6d1f170b94d0f33f61c0e0765c
[448]:
  https://github.com/da-moon/podinfo/commit/6aeea05f1d9907a078b5e5a731686a6ba11739cc
[449]:
  https://github.com/da-moon/podinfo/commit/82fea7885b1910dfa69f8812e94f87a31ed34b53
[450]:
  https://github.com/da-moon/podinfo/commit/24c7238e8ede1dc982fb5350caf00bda517e12e2
[451]:
  https://github.com/da-moon/podinfo/commit/3d8d786808f23ffbf6d5aeed7ecff15129181166
[452]:
  https://github.com/da-moon/podinfo/commit/7b204416f0a90a32e6c95147a8ab31b7add0ebb4
[453]:
  https://github.com/da-moon/podinfo/commit/b95cfbe6753378115483f481352209a24c9d42e6
[454]:
  https://github.com/da-moon/podinfo/commit/63d43cefa539a66e55c37a6ac99e7f9f088deeb2
[455]:
  https://github.com/da-moon/podinfo/commit/87a96580529be827f2cded9526ce85b8283a6009
[456]:
  https://github.com/da-moon/podinfo/commit/40abb30e6c5e96c3d54d562e7cb0cbb98d01d58f
[457]:
  https://github.com/da-moon/podinfo/commit/eb1bae80d4d2a44c943b4ceacd665355b9f12731
[458]:
  https://github.com/da-moon/podinfo/commit/bee96683fc2123ad520cd0d18ee8629d0f118c1f
[459]:
  https://github.com/da-moon/podinfo/commit/5bb284d42a6a6920151e01e4a33483329323d730
[460]:
  https://github.com/da-moon/podinfo/commit/640d343372274c71dae0b59724a0ba424e3f9064
[461]:
  https://github.com/da-moon/podinfo/commit/aabfa50a668fc524e7fd442e3dfa8118be08656f
[462]:
  https://github.com/da-moon/podinfo/commit/8c886060c9cdc8a5f55f0c4c1a967d53b8fb394a
[463]:
  https://github.com/da-moon/podinfo/commit/4f3450265b87b850bafa75890e78b148ab72c828
