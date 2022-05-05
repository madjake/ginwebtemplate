#!/bin/sh

# This will build the IFF binary. By default it will build one for the current operating system
# and architecture. This accepts two arguments that will map to the Operating System
# and Architecture you wish to cross compile for.

# EXAMPLES:
# ./bin/build.sh linux amd64
# ./bin/build.sh linux 386
# ./bin/build.sh windows amd64
# ./bin/build.sh darwin 386
#
# See docs for all valid values: https://golang.org/doc/install/source#environment
cd "$(dirname "$0")"

GOOS=$1
GOARCH=$2

if [ -z "$GOOS" ] || [ -z "$GOARCH" ]; then
  go build -o webserver ../.
else
  go build -o webserver-$GOOS-$GOARCH ../.
fi