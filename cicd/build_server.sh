#! /bin/sh

# Exit immediately if a command exits with a non-zero status
set -e

# This path should be GOPATH=/go/src/github.com/.... when using the golang container
export GOPATH=$(pwd)/go
echo <--NEW GOPATH-->
go env
echo <--CURRENT DIRECTORY-->
pwd
echo <--Nested Folders-->
ls -aR
echo <--GO BUILD COMMAND-->
# cd /go/src/github.com/JosephZoeller/maritime-royale

go build ./go/src/github.com/JospehZoeller/maritime-royale/cmd/server