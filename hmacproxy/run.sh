#! /bin/bash -e

cd /usr/local/18f/hmacproxy

LOGS=/var/log/hmacproxy

if [ ! -d $LOGS ]; then
  mkdir -p $LOGS
fi

source /etc/profile.d/gvm.sh
source config/env-secret.sh

HMACPROXY_SERVER_PORT=8083
HMACPROXY_PROXY_PORT=8084
HMACPROXY_SIGN_HEADER="Team-Api-Signature"
HMACPROXY_HEADERS="Content-Type,Date"

if [ "$1" = "run-server" ]; then
  nohup $GOPATH/bin/hmacproxy -auth -port $HMACPROXY_SERVER_PORT \
    -secret $HMACAUTH_SECRET -sign-header $HMACPROXY_SIGN_HEADER \
    -headers $HMACPROXY_HEADERS \
    >>$LOGS/access.log 2>&1 &
fi
if [ "$1" = "run-proxy" ]; then
  UPSTREAM="$2"

  if [ -z "$UPSTREAM" ]; then
    echo "No upstream server specified"
    exit 1
  fi

  nohup $GOPATH/bin/hmacproxy -port $HMACPROXY_PROXY_PORT \
    -secret $HMACAUTH_SECRET -sign-header $HMACPROXY_SIGN_HEADER \
    -headers $HMACPROXY_HEADERS -upstream $UPSTREAM \
    >>$LOGS/access.log 2>&1 &
fi
