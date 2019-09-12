#!/usr/bin/env bash

# usage:
# ./run_tests.sh                         # local, go 1.11
# ./run_tests.sh docker                  # docker, go 1.11
# ./run_tests.sh podman                  # podman, go 1.11

set -ex

# The script does automatic checking on a Go package and its sub-packages,
# including:
# 1. gofmt         (http://golang.org/cmd/gofmt/)
# 2. go vet        (http://golang.org/cmd/vet)
# 3. gosimple      (https://github.com/dominikh/go-simple)
# 4. unconvert     (https://github.com/mdempsky/unconvert)
# 5. ineffassign   (https://github.com/gordonklaus/ineffassign)
# 6. race detector (http://blog.golang.org/race-detector)

# golangci-lint (github.com/golangci/golangci-lint) is used to run each each
# static checker.

# Default GOVERSION
[[ ! "$GOVERSION" ]] && GOVERSION=1.11
REPO=pfcdata

testrepo () {
  TMPDIR=$(mktemp -d)
  TMPFILE=$(mktemp)
  export GO111MODULE=on

  go version

  ROOTPATH=$(go list -m -f {{.Dir}} 2>/dev/null)
  ROOTPATHPATTERN=$(echo $ROOTPATH | sed 's/\\/\\\\/g' | sed 's/\//\\\//g')
  MODPATHS=$(go list -m -f {{.Dir}} all 2>/dev/null | grep "^$ROOTPATHPATTERN")

  # Test application install
  go build
  (cd cmd/rebuilddb && go build)
  (cd cmd/rebuilddb2 && go build)
  (cd cmd/scanblocks && go build)

  # Check tests
  git clone https://github.com/pfclabs/bug-free-happiness $TMPDIR/test-data-repo
  tar xvf $TMPDIR/test-data-repo/stakedb/test_ticket_pool.bdgr.tar.xz -C ./stakedb

  # run tests on all modules
  for MODPATH in $MODPATHS; do
    env GORACE='halt_on_error=1' go test -v -race $(cd $MODPATH && go list -m)
  done

  # check linters
  golangci-lint run --deadline=10m \
    --disable-all \
    --enable govet \
    --enable staticcheck \
    --enable gosimple \
    --enable unconvert \
    --enable ineffassign \
    --enable structcheck \
    --enable goimports \
    --enable misspell \
    --enable unparam

  # webpack
  npm install
  npm run build

  echo "------------------------------------------"
  echo "Tests completed successfully!"

  # Remove all the tests data
  rm -rf $TMPDIR $TMPFILE
}

testrepo