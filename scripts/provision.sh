#!/bin/bash

apt-get -y install golang git
mkdir -p /home/vagrant/.go
chown -R 1000:1000 /home/vagrant/.go
echo "export GOROOT=/usr/lib/go-1.6" >> /home/vagrant/.profile
echo "export GOPATH=/home/vagrant/.go" >> /home/vagrant/.profile
