#! /bin/bash

. /etc/profile.d/gvm.sh
go get github.com/bitly/oauth2_proxy && strip $GOPATH/bin/oauth2_proxy
