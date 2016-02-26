#! /bin/bash -e

source /etc/profile.d/gvm.sh
source config/env-secret.sh

HMACPROXY_SERVER_PORT=8083
HMACPROXY_PROXY_PORT=8084
HMACPROXY_SIGN_HEADER="Team-Api-Signature"
HMACPROXY_HEADERS="Content-Type,Date"

if [ "$1" = "run-server" ]; then
  exec hmacproxy -auth -port $HMACPROXY_SERVER_PORT -secret $HMACAUTH_SECRET \
    -sign-header $HMACPROXY_SIGN_HEADER -headers $HMACPROXY_HEADERS
fi
if [ "$1" = "run-proxy" ]; then
  UPSTREAM="$2"

  if [ -z "$UPSTREAM" ]; then
    echo "No upstream server specified"
    exit 1
  fi

  exec hmacproxy -port $HMACPROXY_PROXY_PORT -secret $HMACAUTH_SECRET \
    -sign-header $HMACPROXY_SIGN_HEADER -headers $HMACPROXY_HEADERS \
    -upstream $UPSTREAM
fi
exec "$@"
