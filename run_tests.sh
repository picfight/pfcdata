set -ex

# Default GOVERSION
[[ ! "$GOVERSION" ]] && GOVERSION=1.11
REPO=pfcdata

testrepo () {
  GO=go
  if [[ $GOVERSION == 1.10 ]]; then
    GO=vgo
  fi

  $GO version

  # binary needed for RPC tests
  env CC=gcc $GO build
  cp "$REPO" "$GOPATH/bin/"

  # run tests on all modules
  ROOTPATH=$($GO list -m -f {{.Dir}} 2>/dev/null)
  ROOTPATHPATTERN=$(echo $ROOTPATH | sed 's/\\/\\\\/g' | sed 's/\//\\\//g')
  MODPATHS=$($GO list -m -f {{.Dir}} all 2>/dev/null | grep "^$ROOTPATHPATTERN"\
    | sed -e "s/^$ROOTPATHPATTERN//" -e 's/^\\//' -e 's/^\///')
  MODPATHS=". $MODPATHS"
  for module in $MODPATHS; do
    echo "==> ${module}"
    (cd $module && env GORACE='halt_on_error=1' CC=gcc $GO test -short -race \
	  -tags rpctest ./...)
  done

  # check linters
  if [[ $GOVERSION != 1.10 ]]; then
    # linters do not work with modules yet
    go mod vendor
    unset GO111MODULE

  fi

  echo "------------------------------------------"
  echo "Tests completed successfully!"
}

DOCKER=
[[ "$1" == "docker" || "$1" == "podman" ]] && DOCKER=$1
if [ ! "$DOCKER" ]; then
    testrepo
    exit
fi

# use Travis cache with docker
DOCKER_IMAGE_TAG=picfight-golang-builder-$GOVERSION
mkdir -p ~/.cache
if [ -f ~/.cache/$DOCKER_IMAGE_TAG.tar ]; then
  # load via cache
  $DOCKER load -i ~/.cache/$DOCKER_IMAGE_TAG.tar
else
  # pull and save image to cache
  $DOCKER pull picfight/$DOCKER_IMAGE_TAG
  $DOCKER save picfight/$DOCKER_IMAGE_TAG > ~/.cache/$DOCKER_IMAGE_TAG.tar
fi

$DOCKER run --rm -it -v $(pwd):/src:Z picfight/$DOCKER_IMAGE_TAG /bin/bash -c "\
  rsync -ra --filter=':- .gitignore'  \
  /src/ /go/src/github.com/picfight/$REPO/ && \
  dir && \
  env GOVERSION=$GOVERSION GO111MODULE=on bash run_tests.sh"
