#!/bin/bash

mkdir -p bin
vagrant up
vagrant ssh -c "pushd /opt/zfsnap
go get github.com/jarosser06/go-zsnap/zsnap
go build -o bin/zfsnap zfsnap.go
popd"
