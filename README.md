# pfcdata

[![Build Status](https://img.shields.io/travis/picfight/pfcdata.svg)](https://travis-ci.org/picfight/pfcdata)
[![Latest tag](https://img.shields.io/github/tag/picfight/pfcdata.svg)](https://github.com/picfight/pfcdata/tags)
[![Go Report Card](https://goreportcard.com/badge/github.com/picfight/pfcdata)](https://goreportcard.com/report/github.com/picfight/pfcdata)
[![ISC License](https://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)

## Requirements

- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [Go](http://golang.org/doc/install) 1.13+
- Running `pfcd`-node synchronized to the current best block on the
  network.
- (Optional) PostgreSQL 9.6+, if running in "full" mode. v10.x is recommended
  for improved dump/restore formats and utilities.

## Setup

The following instructions assume a Unix-like shell (e.g. bash).

- Verify Go installation:

      go env GOROOT GOPATH

- Ensure `$GOPATH/bin` is on your `$PATH`.

- Clone the pfcdata repository. It is conventional to put it under `GOPATH`, but
  this is no longer necessary with go module.

      git clone https://github.com/picfight/pfcdata $GOPATH/src/github.com/picfight/pfcdata

- Install a C compiler. 

```bash
cd $GOPATH/src/github.com/picfight/pfcdata
set GO111MODULE=on

go version
go clean -testcache
go build -v ./...
go test -v ./...
go install
```

## License

This project is licensed under the ISC License. See the [LICENSE](LICENSE) file for details.
