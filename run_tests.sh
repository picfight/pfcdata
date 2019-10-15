#!/usr/bin/env bash
set -ex

GO111MODULE=on

  go version
  go clean -testcache
  go build -v ./...
  go test -v ./...
  go install
