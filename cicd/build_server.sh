#! /bin/sh

# Exit immediately if a command exits with a non-zero status
# set -e

# This path should be GOPATH=/go/src/github.com/.... when using the golang container
export GOPATH=$PWD/gopath
export PATH=$PWD/gopath/bin:$PATH

echo NEW GOPATH
go env
echo CURRENT DIRECTORY
pwd
echo Nested Folders
ls -aR
echo GO BUILD COMMAND
# cd /go/src/github.com/JosephZoeller/maritime-royale
go get -v -u github.com/JosephZoeller/maritime-royale
cd gopath/src/github.com/JosephZoeller/maritime-royale/

echo Directory after cd
pwd
go build ./cmd/server/server.go
ls