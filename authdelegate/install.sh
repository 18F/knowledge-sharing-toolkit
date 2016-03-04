#! /bin/bash

. /etc/profile.d/gvm.sh
go get github.com/18F/authdelegate && strip $GOPATH/bin/authdelegate
