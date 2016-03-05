#! /bin/bash

. /etc/profile.d/gvm.sh
go get github.com/18F/hmacproxy && strip $GOPATH/bin/hmacproxy
