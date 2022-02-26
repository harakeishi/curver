![logo](logo.png)
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->
[![GitHub release](https://img.shields.io/github/release/harakeishi/curver.svg)](https://github.com/harakeishi/curver/releases) [![Go Report Card](https://goreportcard.com/badge/github.com/harakeishi/curver)](https://goreportcard.com/report/github.com/harakeishi/curver) 

`curver` is a simple way to display the version of a CUI tool made with go.
`curver` was named as an abbreviation for `current version`.

### Table of Contents
- [Installation](#installation)
- [Importing](#importing)
- [Documentation](#documentation)
- [usage](#usage)
  - [use go build](#use-go-build)
  - [use goreleaser](#use-goreleaser)
  - [If you only want the version](#if-you-only-want-the-version)
- [License](#license)
- [Contributors ✨](#contributors-)

## Installation
```bash
go get github.com/harakeishi/curver
```

## Importing
```go
import (
    "github.com/harakeishi/curver"
)
```

## Documentation
Visit the docs on [GoDoc](https://pkg.go.dev/github.com/harakeishi/curver)

## usage
If the value using ldflag is stored in the variable Version, that value will be displayed.
Otherwise, it will display the build information embedded in the running binary.

### use go build
```go:main.go
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
$ go build -ldflags "-X github.com/harakeishi/curver.Version=$(git describe --tags)" -o ./main
```
### use goreleaser
If you are using [goreleaser](https://goreleaser.com/) to do the release, do the following

```yml
builds:
  - eldflags:
      - -s -w -X github.com/harakeishi/curver.Version={{.Version}}

```

### If you only want the version
The following will return the version as a string.

```go:main.go
package main

import (
	"github.com/harakeishi/curver"
)

func main () {
    version := curver.GetVersion()
}
```

## License
Copyright (c) 2022 harakeishi
Licensed under [MIT](LICENSE)

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://yaserarenai.com"><img src="https://avatars.githubusercontent.com/u/44335168?v=4?s=100" width="100px;" alt=""/><br /><sub><b>原　慧士</b></sub></a><br /><a href="https://github.com/harakeishi/curver/commits?author=harakeishi" title="Code">💻</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!