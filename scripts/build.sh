#!/bin/bash

mkdir -p bin
vagrant up
vagrant ssh -c "pushd /opt/zfsnap &> /dev/null
go get github.com/TAMUArch/go-zsnap/zsnap
go build -o bin/zfsnap zfsnap.go
popd &> /dev/null"
