![logo](logo.png)
[![GitHub release](https://img.shields.io/github/release/harakeishi/curver.svg)](https://github.com/harakeishi/curver/releases) [![Go Report Card](https://goreportcard.com/badge/github.com/harakeishi/curver)](https://goreportcard.com/report/github.com/harakeishi/curver) [![Test](https://github.com/harakeishi/curver/actions/workflows/test.yml/badge.svg)](https://github.com/harakeishi/curver/actions/workflows/test.yml)

`curver` is a simple way to display the version of a CUI tool made with go.
`curver` was named as an abbreviation for `current version`.

## Table of Contents
- [Installation](#installation)
- [Importing](#importing)
- [Documentation](#documentation)
- [usage](#usage)
  - [normal](#normal)
  - [use goreleaser](#use-goreleaser)
- [License](#license)

# Installation
```bash
go get github.com/harakeishi/curver
```

# Importing
```go
import (
    "github.com/harakeishi/curver"
)
```

# Documentation
Visit the docs on [GoDoc](https://pkg.go.dev/github.com/harakeishi/curver)

# usage
If the value using ldflag is stored in the variable Version, that value will be displayed.
Otherwise, it will display the build information embedded in the running binary.

## normal
```go
// main.go
package main

import (
	"github.com/harakeishi/curver"
)

func main () {
    curver.EchoVersion()
}
```

``` bash
$ go build -ldflags '-X github.com/harakeishi/curver.Version=v0.1.0' -o ./main
```

This will display the following format.

```bash
$ ./main
version: v0.1.0
```

If you want to embed the result of 'git tag', you can do the following

```
$ go build -ldflags "-X github.com/harakeishi/curver.Version=$(git describe --tags)" -o ./cmds 
```
## use goreleaser
If you are using [goreleaser](https://goreleaser.com/) to do the release, do the following

```yml
builds:
  - eldflags:
      - -s -w -X github.com/harakeishi/curver.Version={{.Version}}

```

# License
Copyright (c) 2022 harakeishi
Licensed under [MIT](LICENSE)
