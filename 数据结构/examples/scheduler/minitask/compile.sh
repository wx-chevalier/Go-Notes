#!/bin/bash 

dir=$(cd -P -- "$(dirname -- "$0")" && pwd -P)

export GOROOT=/usr/local/go
export PATH=$PATH:/usr/local/go/bin
export GOBIN=$dir/bin
export GOPATH=$GOPATH:$dir
export PKG_CONFIG_PATH=/usr/lib/pkgconfig/

echo "/usr/local/go/bin/go install "$dir"/src/"$1
go install $dir/src/$1