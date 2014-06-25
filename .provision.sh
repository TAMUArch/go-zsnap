#!/bin/bash

GOVERSION=1.3


pushd /tmp
wget http://public-mirror.tamu.edu/golang/go${GOVERSION}.omnios-amd64.tar.gz
tar -xf go${GOVERSION}.omnios-amd64.tar.gz -C /opt
popd
mkdir -p /export/home/vagrant/.go
chown -R 100:100 /export/home/vagrant/.go
echo "export PATH=\$PATH:/opt/go/bin" >> /export/home/vagrant/.profile
echo "export GOROOT=/opt/go" >> /export/home/vagrant/.profile
echo "export GOPATH=/export/home/vagrant/.go" >> /export/home/vagrant/.profile
pkg install --accept developer/versioning/git
