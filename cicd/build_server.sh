#!bin/sh

# Exit immediately if a command exits with a non-zero status
set -e

# This path should be GOPATH=/go/src/github.com/.... when using the golang container
cd /go/src/github.com/JosephZoeller/maritime-royale

go build ./cmd/server