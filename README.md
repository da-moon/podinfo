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

```console
❯ podinfo server -h
2023/02/14 21:48:20 profile: cpu profiling enabled, cpu.pprof
Usage: podinfo server [options]

Starts podinfo server.

Server Options:

  -api-addr=<value>
     flag used to set the address podinfo is listening on.
     This can also be specified via the 'PODINFO_API_ADDR' env variable.

  -dev
     flag used to start the server in development mode
     This can also be specified via the 'PODINFO_DEVEL' env variable
     (true|false)

  -log-level=<value>
     flag used to set stdlogger level.
     This can also be specified via the 'PODINFO_LOG_LEVEL' env
     variable.

  -node-name=<value>
     flag used to set podinfo node name.
     This can also be specified via the 'PODINFO_NODE_NAME' env
     variable.

 Options:

  -redis-addr=<value>
     flag used to set  address
     it is in host:port address format.
     This can also be specified via the 'PODINFO_REDIS_ADDR' env
     variable

  -redis-client-name=<CLIENT SETNAME ClientName>
     flag used to set  client name.
     ClientName will execute the `CLIENT SETNAME ClientName` command for
     each conn.
     This can also be specified via the 'PODINFO_REDIS_CLIENT_NAME' env
     variable

  -redis-conn-max-idle-time=<value>
     flag used to set the maximum amount of time a connection may be
     idle.Should be less than server's timeout
     Expired connections may be closed lazily before reuseIf d <= 0,
     connections are not closed due to a connection's idle time.
     Default is 30 minutes. -1 disables idle timeout check.
     This can also be specified via the
     'PODINFO_REDIS_CONN_MAX_IDLE_TIME' env variable

  -redis-conn-max-lifetime=<value>
     flag used to set the maximum amount of time a connection may be
     reused.
     Expired connections may be closed lazily before reuse.If <= 0,
     connections are not closed due to a connection's age.
     Default is to not close idle connections.
     This can also be specified via the
     'PODINFO_REDIS_CONN_MAX_LIFETIME' env variable

  -redis-context-timeout-enabled
     flag used to set  context timeout enabled.
     contextTimeoutEnabled controls whether the client respects context
     timeouts and deadlines.
     See https://.uptrace.dev/guide/go--debugging.html#timeouts
     This can also be specified via the
     'PODINFO_REDIS_CONTEXT_TIMEOUT_ENABLED' env variable

  -redis-db=<value>
     flag used to set  database.
     Database to be selected after connecting to the server.
     This can also be specified via the 'PODINFO_REDIS_db' env variable

  -redis-dial-timeout=<0>
     flag used to set timeout for socket reads.
     If reached, commands will failwith a timeout instead of blocking.
     Supported values:
     - `0` - default timeout (3 seconds).
     - `-1` - no timeout (block indefinitely).
     - `-2` - disables SetReadDeadline calls completely.
     Default is 5 seconds.
     This can also be specified via the 'PODINFO_REDIS_DIAL_TIMEOUT' env
     variable

  -redis-max-idle-conns=<value>
     flag used to set the maximum number of idle connections.
     This can also be specified via the 'PODINFO_REDIS_MAX_IDLE_CONNS'
     env variable

  -redis-max-retries=<value>
     flag used to set the maximum number of retries before giving up.
     Default is 3 retries; -1 disables retries.
     This can also be specified via the 'PODINFO_REDIS_MAX_RETRIES' env
     variable

  -redis-max-retry-backoff=<value>
     flag used to set Specifies the maximum backoff between each retry.
     Default is 512 milliseconds; -1 disables backoff.
     This can also be specified via the
     'PODINFO_REDIS_MAX_RETRY_BACKOFF' env variable

  -redis-min-idle-conns=<value>
     flag used to set the minimum number of idle connections.
     it is useful when establishing new connection is slow.
     This can also be specified via the 'PODINFO_REDIS_MIN_IDLE_CONNS'
     env variable

  -redis-min-retry-backoff=<value>
     flag used to set minimum backoff between each retry.
     Default is 8 milliseconds; -1 disables backoff.
     This can also be specified via the
     'PODINFO_REDIS_MIN_RETRY_BACKOFF' env variable

  -redis-password=<value>
     flag used to set  password.
     Optional password. Must match the password specified in
     therequirepass server configuration option (if connecting to a  5.0
     instance, or lower),or the User password when connecting to a  6.0
     instance, or greater,that is using the  ACL system.
     This can also be specified via the 'PODINFO_REDIS_PASSWORD' env
     variable

  -redis-pool-fifo
     flag used to set  pool fifo.
     Type of connection pool.
     true for FIFO pool, false for LIFO pool.
     Note that FIFO has slightly higher overhead compared to LIFO,but it
     helps closing idle connections faster reducing the pool size.
     This can also be specified via the 'PODINFO_REDIS_POOL_FIFO' env
     variable

  -redis-pool-size=<value>
     flag used to set  pool size.
     Maximum number of socket connections.
     Default is 10 connections per every available CPU
     This can also be specified via the 'PODINFO_REDIS_POOL_SIZE' env
     variable

  -redis-pool-timeout=<value>
     flag used to set  pool timeout.
     Amount of time client waits for connection if all connectionsare
     busy before returning an error.
     Default is readTimeout + 1 second.
     This can also be specified via the 'PODINFO_REDIS_POOL_TIMEOUT' env
     variable

  -redis-read-timeout=<value>
     flag used to set  read timeout.
     Timeout for socket reads. If reached, commands will failwith a
     timeout instead of blocking. Supported values:
     - '0' - default timeout (3 seconds).
     - '-1' - no timeout (block indefinitely).
     - '-2' - disables SetReadDeadline calls completely.
     This can also be specified via the 'PODINFO_REDIS_READ_TIMEOUT' env
     variable

  -redis-username=<value>
     flag used to set  username.
     Use the specified username to authenticate the current
     connectionwith one of the connections defined in the ACL list when
     connectingto a  6.0 instance, or greater, that is using the  ACL
     system.
     This can also be specified via the 'PODINFO_REDIS_USERNAME' env
     variable

  -redis-write-timeout=<0>
     flag used to set  write timeout.
     Timeout for socket writes. If reached, commands will failwith a
     timeout instead of blocking.  Supported values:
     - `0` - default timeout (3 seconds).
     - `-1` - no timeout (block indefinitely).
     - `-2` - disables SetWriteDeadline calls completely.
     This can also be specified via the 'PODINFO_REDIS_WRITE_TIMEOUT'
     env variable

Telemetry Options:

  -metrics-prefix=<value>
     flag used to set default metrics prefix.
     This can also be specified via the 'PODINFO_METRICS_PREFIX' env
     variable

  -prometheus-retention-time=<value>
     flag used to set prometheus retention time.
     This can also be specified via the
     'PODINFO_PROMETHEUS_RETENTION_TIME' env variable

  -statsd-addr=<value>
     flag used to set statsd address
     This can also be specified via the 'STATSD_ADDR' env variable

  -statsite-addr=<value>
     flag used to set statsite address.
     This can also be specified via the 'STATSITE_ADDR' env variable
```

## Code Statistics

```console
===============================================================================
 Language            Files        Lines         Code     Comments       Blanks
===============================================================================
 Dockerfile              2          657          567           89            1
 Go                    243        17554        13492         2553         1509
 HCL                     1          129           49           80            0
 JavaScript              1           25           23            2            0
 JSON                    5          291          291            0            0
 Shell                   5          205          179           22            4
 Plain Text              2          193            0          193            0
 TOML                    1           57           57            0            0
 YAML                   10          943          864           71            8
-------------------------------------------------------------------------------
 Markdown                3         1395            0         1138          257
 |- BASH                 1           25           22            1            2
 (Total)                           1420           22         1139          259
===============================================================================
 Total                 273        21449        15522         4148         1779
===============================================================================
```

## Usage and Demo

- Build The Binary

```bash
❯ just build
                      _   _            __
  _ __     ___     __| | (_)  _ __    / _|   ___
 | '_ \   / _ \   / _` | | | | '_ \  | |_   / _ \
 | |_) | | (_) | | (_| | | | | | | | |  _| | (_) |
 | .__/   \___/   \__,_| |_| |_| |_| |_|    \___/
 |_|

# Build Command ------------------------------------------

go build \
  -ldflags \
  '
  -X "github.com/da-moon/northern-labs-interview/build/go/version.BuildDate=02/14/23"
  -X "github.com/da-moon/northern-labs-interview/build/go/version.BuildUser=gitpod"
  -X "github.com/da-moon/northern-labs-interview/build/go/version.Branch=master"
  -X "github.com/da-moon/northern-labs-interview/build/go/version.Revision=68e97799e8952aa04bd0d692943c494377c314c6"
  ' -o /workspace/northern-labs-interview/bin/podinfo  /workspace/northern-labs-interview/cmd/podinfo
```

In case the build fails, due to missing dependencies, you can run
`just bootstrap`. It installs all necessary dependencies and configures the
environment on `Debian` , `Alpine` and `Arch` Linux distributions.

You can also use the `gitpod` environment either by directly opening it in
gitpod, or building the gitpod Docker image (`.gp/Dockerfile`) and running
commands inside it. There is a target for building the `gitpod` image in
`docker-bake.hcl` file so it should be relatively easy to build it; the biggest
issue is that the docker image is fairly large so it can take a while to build
it; the other issue is that it is based on `archlinux` image which does not
support `aarch64` architecture so as an example, you cannot run it on Mac's
with `M1` chip.

- An alternative to building the binary is using the docker image
  `fjolsvin/podinfo`

```bash
docker run --rm -it fjolsvin/podinfo:1.0.0
```

I prefer building against `master` as docker images are only pushed on tags and
the repo is under heavy development so the image might not include the latest
changes.

- start redis server with docker compose:

```bash
docker compose up redis -d
```

- start the server by running
  `PODINFO_REDIS_CLIENT_NAME="$(whoami)" bin/podinfo server -log-level=trace -redis-password="foobared"`

```console
✦ ❯ PODINFO_REDIS_CLIENT_NAME="$(whoami)" bin/podinfo server -log-level=trace -redis-password="foobared"
2023/02/15 16:23:24 profile: cpu profiling enabled, cpu.pprof
                                                              ██████   ██████  ██████  ██ ███    ██ ███████  ██████
                                                              ██   ██ ██    ██ ██   ██ ██ ████   ██ ██      ██    ██
                                                              ██████  ██    ██ ██   ██ ██ ██ ██  ██ █████   ██    ██
                                                              ██      ██    ██ ██   ██ ██ ██  ██ ██ ██      ██    ██
                                                              ██       ██████  ██████  ██ ██   ████ ██       ██████



                                                                               INFO  podinfo running!



build info:

                   Version Info: '(branch=master, revision=8ead18d61f80ef2ec763c807a58b440e70b25375)'
                   Build Context: '(go=1.20, user=damoon, date=02/15/23)'

Node info:

                   Log Level: 'TRACE'
                   Development Mode: 'false'
                   Node name: 'archlinux'
                   API addr: '0.0.0.0:2048'

Redis Info:

                   Address: '0.0.0.0:6379'
                   Client Name: 'damoon'
                   DB: '0'
                   MaxRetries: '3'
                   MinRetryBackoff: '0s'
                   MaximumRetryBackoff: '0s'
                   DialTimeout: '0s'
                   ReadTimeout: '0s'
                   WriteTimeout: '0s'
                   PoolFIFO: 'false'
                   PoolSize: '80'
                   PoolTimeout: '0s'
                   MinIdleConns: '0'
                   MaxIdleConns: '0'
                   ConnMaxIdleTime: '30m0s'
                   ConnMaxLifetime: '-1ns'


Telemetry Info:

                   MetricsPrefix: 'podinfo_api'
                   PrometheusRetentionTime: '1m0s'

Log data will now stream in as it occurs:

2023/02/15 16:23:24 [ INFO  ] restful-server successfully bound to host port
2023/02/15 16:23:24 [ INFO  ] Adding log middleware for 'debug-index' handler at '/pprof'
2023/02/15 16:23:24 [ INFO  ] Adding log middleware for 'debug-allocs' handler at '/allocs'
2023/02/15 16:23:24 [ INFO  ] Adding log middleware for 'debug-block' handler at '/block'
2023/02/15 16:23:24 [ INFO  ] Adding log middleware for 'debug-cmdline' handler at '/cmdline'
2023/02/15 16:23:24 [ INFO  ] Adding log middleware for 'debug-goroutine' handler at '/goroutine'
2023/02/15 16:23:24 [ INFO  ] Adding log middleware for 'debug-heap' handler at '/heap'
2023/02/15 16:23:24 [ INFO  ] Adding log middleware for 'debug-mutex' handler at '/mutex'
2023/02/15 16:23:24 [ INFO  ] Adding log middleware for 'debug-profile' handler at '/profile'
2023/02/15 16:23:24 [ INFO  ] Adding log middleware for 'debug-threadcreate' handler at '/threadcreate'
2023/02/15 16:23:24 [ INFO  ] Adding log middleware for 'debug-symbol' handler at '/symbol'
2023/02/15 16:23:24 [ INFO  ] Adding log middleware for 'debug-trace' handler at '/trace'
2023/02/15 16:23:24 [ INFO  ] metrics middleware path = [ /healthz ]  label = [ healthz ]
2023/02/15 16:23:24 [ INFO  ] metrics middleware path = [ /readyz ]  label = [ readyz ]
2023/02/15 16:23:24 [ INFO  ] metrics middleware path = [ /enable ]  label = [ enable ]
2023/02/15 16:23:24 [ INFO  ] metrics middleware path = [ /disable ]  label = [ disable ]
2023/02/15 16:23:24 [ INFO  ] metrics middleware path = [ /env ]  label = [ env ]
2023/02/15 16:23:24 [ INFO  ] metrics middleware path = [ /headers ]  label = [ headers ]
2023/02/15 16:23:24 [ INFO  ] metrics middleware path = [ /delay/{seconds} ]  label = [ delay_{seconds} ]
2023/02/15 16:23:24 [ INFO  ] metrics middleware path = [ /{key} ]  label = [ {key} ]
2023/02/15 16:23:24 [ INFO  ] metrics middleware path = [ /{key} ]  label = [ {key} ]
2023/02/15 16:23:24 [ INFO  ] metrics middleware path = [ /{key} ]  label = [ {key} ]
2023/02/15 16:23:24 [ INFO  ] metrics middleware path = [ /{key} ]  label = [ {key} ]
2023/02/15 16:23:24 [ INFO  ] initializing telementry
2023/02/15 16:23:24 [ INFO  ] metrics : validating configuration
2023/02/15 16:23:24 [ INFO  ] metrics : validating 'MetricsPrefix' configuration value
2023/02/15 16:23:24 [ INFO  ] metrics : 'MetricsPrefix' configuration value was successfully validated
2023/02/15 16:23:24 [ INFO  ] metrics : validating 'StatsiteAddr' configuration value
2023/02/15 16:23:24 [ INFO  ] metrics : 'StatsiteAddr' configuration value was successfully validated
2023/02/15 16:23:24 [ INFO  ] metrics : validating 'StatsdAddr' configuration value
2023/02/15 16:23:24 [ INFO  ] metrics : 'StatsdAddr' configuration value was successfully validated
2023/02/15 16:23:24 [ INFO  ] metrics : validating 'PrometheusRetentionTime' configuration value
2023/02/15 16:23:24 [ INFO  ] metrics : 'PrometheusRetentionTime' configuration value was successfully validated
2023/02/15 16:23:24 [ INFO  ] metrics : validating configuration
2023/02/15 16:23:24 [ INFO  ] metrics : validating 'MetricsPrefix' configuration value
2023/02/15 16:23:24 [ INFO  ] metrics : 'MetricsPrefix' configuration value was successfully validated
2023/02/15 16:23:24 [ INFO  ] metrics : validating 'StatsiteAddr' configuration value
2023/02/15 16:23:24 [ INFO  ] metrics : 'StatsiteAddr' configuration value was successfully validated
2023/02/15 16:23:24 [ INFO  ] metrics : validating 'StatsdAddr' configuration value
2023/02/15 16:23:24 [ INFO  ] metrics : 'StatsdAddr' configuration value was successfully validated
2023/02/15 16:23:24 [ INFO  ] metrics : validating 'PrometheusRetentionTime' configuration value
2023/02/15 16:23:24 [ INFO  ] metrics : 'PrometheusRetentionTime' configuration value was successfully validated
2023/02/15 16:23:24 [ INFO  ] Setting version gauge
2023/02/15 16:23:24 [ INFO  ] Starting prometheus Metrics collector core engine
2023/02/15 16:23:24 [ INFO  ] prometheus Metrics collector core engine successfully initialized
2023/02/15 16:23:24 [ INFO  ] metrics exporter route was successfully initialized.

+ Listening On  : 127.0.0.1:2048
+ API Version (Prefix) : /readyz
+ Routes:
[ GET ] disable-kubernetes-readiness-probe      127.0.0.1:2048/readyz/disable
[ GET ] enable-kubernetes-readiness-probe       127.0.0.1:2048/readyz/enable


+ API Version (Prefix) : /cache
+ Routes:
[ DELETE ] delete       127.0.0.1:2048/cache/{key}
[ GET ] get     127.0.0.1:2048/cache/{key}
[ PUT ] put     127.0.0.1:2048/cache/{key}
[ POST ] post   127.0.0.1:2048/cache/{key}


+ API Version (Prefix) : /debug
+ Routes:
[ GET ] debug-allocs    127.0.0.1:2048/debug/allocs
[ GET ] debug-block     127.0.0.1:2048/debug/block
[ GET ] debug-cmdline   127.0.0.1:2048/debug/cmdline
[ GET ] debug-goroutine 127.0.0.1:2048/debug/goroutine
[ GET ] debug-heap      127.0.0.1:2048/debug/heap
[ GET ] debug-mutex     127.0.0.1:2048/debug/mutex
[ GET ] debug-index     127.0.0.1:2048/debug/pprof
[ GET ] debug-profile   127.0.0.1:2048/debug/profile
[ GET ] debug-symbol    127.0.0.1:2048/debug/symbol
[ GET ] debug-threadcreate      127.0.0.1:2048/debug/threadcreate
[ GET ] debug-trace     127.0.0.1:2048/debug/trace


+ Global API :
+ Routes:
[ GET ] simulate-delay  127.0.0.1:2048/delay/{seconds}
[ GET ] get-environment-variables       127.0.0.1:2048/env
[ GET ] get-request-headers     127.0.0.1:2048/headers
[ GET ] kubernetes-liveness-probe       127.0.0.1:2048/healthz
[ GET ] metrics 127.0.0.1:2048/metrics
[ GET ] kubernetes-readiness-probe      127.0.0.1:2048/readyz

2023/02/15 16:23:24 [ INFO  ] restful-server initializing NotFound route handler
2023/02/15 16:23:24 [ INFO  ] restful-server routers are ready to serve client requests
2023/02/15 16:23:24 [ INFO  ] asynchronous API endpoint initialization started
```

- There are `just` recipes for testing various API endpoints

```console
❯ just -l | grep probe
    delay-probe             # send a GET API request to /delay/{seconds} endpoint
    env-probe               # send a GET API request to /env endpoint
    headers-probe           # send a GET API request to /headers endpoint
    liveness-probe          # send a GET API request to /healthz endpoint
    readiness-probe         # send a GET API request to /readyz endpoint
    readiness-probe-disable # send a GET API request to /readyz/disable endpoint
    disable-readiness-probe # alias for `readiness-probe-disable`
    readiness-probe-enable  # send a GET API request to /readyz/enable endpoint
    enable-readiness-probe  # alias for `readiness-probe-enable`
```

In case you are wondering what command is getting executed in the recipe, you
can use the `--dry-run` flag. e.g :

```console
❯ just --dry-run env-probe
#!/usr/bin/env bash
echo "─── SUCCESS ──────────────────────────────────────────────────────────────────"
URI="env"
VERB="GET"
echo "❯ Sending ${VERB} request to /${URI}"
URL="http://localhost:${PODINFO_SERVER_PORT}/${URI}"
resp="$(curl -o - -sSl --request "${VERB}" "${URL}" )";
echo "${resp}" | jq -r || true
status_code="$(curl -s -o /dev/null -w "%{http_code}" "${URL}" || true)"
echo "Status Code: ${status_code}"
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
✦ ❯ just -l
Available recipes:
    bootstrap               # installs dependencies and prepares development environment
    b                       # alias for `bootstrap`
    bootstrap-bash          # install all bash toolings
    bootstrap-git           # installs necessary git tools and configures git
    bootstrap-go            # install all go toolings
    bootstrap-json
    bootstrap-markdown      # install all markdown toolings
    bootstrap-os-pkgs       # this target installs a collection of core os packages. supports (debian, arch, alpine)
    bootstrap-pre-commit    # ensures tools for making sane commits are installed and initializes pre-commit
    pc                      # alias for `bootstrap-pre-commit`
    bootstrap-semver        # bootstrap semantic versioning toolings
    build-go                # cross-compile go binaries for all supported platforms
    build                   # alias for `build-go`
    cache-delete-probe      # send a DELETE API request to /cache/{key} endpoint
    cache-get-probe         # send a GET API request to /cache/{key} endpoint
    cache-post-probe        # send a POST API request to /cache/{key} endpoint
    cache-put-probe         # send a PUT API request to /cache/{key} endpoint
    clean-go                # removes build binaries (bin/) and tmp/ directory in repo's root
    clean                   # alias for `clean-go`
    commit                  # help guide the developers make conventional commits. it is recommended to use this target to make new commits
    c                       # alias for `commit`
    default                 # `default` target, i.e target execued when just is called without any arguments
    delay-probe             # send a GET API request to /delay/{seconds} endpoint
    env-probe               # send a GET API request to /env endpoint
    format                  # run all formatters
    f                       # alias for `format`
    fmt                     # alias for `format`
    format-bash             # detect and format all bash scripts
    bash-fmt                # alias for `format-bash`
    shfmt                   # alias for `format-bash`
    format-go               # format all go files
    go-fmt                  # alias for `format-go`
    gofmt                   # alias for `format-go`
    format-json             # detect and format all json files
    json-fmt                # alias for `format-json`
    format-just             # format and stage the justfile
    just-fmt                # alias for `format-just`
    generate-changelog      # generate markdown and pdf changelog files
    gc                      # alias for `generate-changelog`
    git-add                 # uses fzf to list git changes and help developers stage them
    ga                      # alias for `git-add`
    git-fetch               # fetches latest changes from upstream and removes any local branches that have been deleted in upstream
    gf                      # alias for `git-fetch`
    headers-probe           # send a GET API request to /headers endpoint
    kary-comments           # adds support for extra languages to Kary Comments VSCode extension
    kc                      # alias for `kary-comments`
    kill                    # send SIGTERM to running binary to stop it
    lint                    # run all linters
    lint-bash               # lint all shellscripts
    shellcheck              # alias for `lint-bash`
    lint-go                 # run golangci-lint with repo specific config
    golangci-lint           # alias for `lint-go`
    liveness-probe          # send a GET API request to /healthz endpoint
    major-release           # generate changelog and create and push a new major release tag
    mar                     # alias for `major-release`
    minor-release           # generate changelog and create and push a new minor release tag
    patch-release           # generate changelog and create and push a new patch release tag
    pr                      # alias for `patch-release`
    readiness-probe         # send a GET API request to /readyz endpoint
    readiness-probe-disable # send a GET API request to /readyz/disable endpoint
    disable-readiness-probe # alias for `readiness-probe-disable`
    readiness-probe-enable  # send a GET API request to /readyz/enable endpoint
    enable-readiness-probe  # alias for `readiness-probe-enable`
    release                 # runs go-releaser (for testing) to build binary(s) and generate a release archive without publishing.
    run                     # build and start the server and forward logs to ./tmp/server/log
    snapshot                # take a tarball 'snapshot' of the repository.
    vscode-tasks            # generate vscode tasks.json file from justfile
    vt                      # alias for `vscode-tasks`
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
  - \[x] Redis Group : There is a minor issue in the code that initializes the
    server which prevents server startup when redis config is not passed. This
    will be addressed immediately
    - \[x] pre-flight redis connection check
    - \[x] **POST** `/cache/{key}`
      - \[x] Implementation
      - \[x] validation
      - \[x] E2E test automation
      - \[x] `swagger` config
    - \[x] **PUT** `/cache/{key}`
      - \[x] Implementation
      - \[x] validation
      - \[x] E2E test automation
      - \[x] `swagger` config
    - \[x] **GET** `/cache/{key}`
      - \[x] Implementation
      - \[x] validation
      - \[x] E2E test automation
      - \[x] `swagger` config
    - \[x] **DELETE** `/cache/{key}`
      - \[x] Implementation
      - \[x] validation
      - \[x] E2E test automation
      - \[x] `swagger` config

- \[x] docker-compose file with Redis

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
