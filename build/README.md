# build

## Table of contents

- [Overview][1]

## Overview

Build system files for this project are stored in this directory:

- `version` : Go library that helps with embedding build information into
  binary with `ldflags` directives.
- `go` : Go files used for building the project. [`mage`][2] is used as the
  build system.
- `just` : Justfiles used for running common tasks.

[1]: #overview
[2]: https://magefile.org
