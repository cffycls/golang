#!/bin/bash
wget -O /tmp/go.tgz "https://dl.google.com/go/go1.14.2.src.tar.gz"
tar -C /usr/local -xzf /tmp/go.tgz  && rm -f /tmp/go.tgz
ln -s /usr/local/go/bin/ /usr/local/bin/

#export GOPATH	工作目录
#export GOBIN 	bin目录
sed -i '$a export GOPATH=/home/golang' ~/.bashrc
sed -i '$a export GOBIN=/usr/local/go/bin' ~/.bashrc
sed -i '$a export PATH=$PATH:$GOPATH:$GOBIN' ~/.bashrc
