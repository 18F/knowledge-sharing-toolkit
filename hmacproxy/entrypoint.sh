#! /bin/bash -e

source /etc/profile.d/gvm.sh
source config/env-secret.sh

if [ "$1" = "run-server" ]; then
  exec hmacproxy -auth -port 8083 -secret $HMACAUTH_SECRET \
    -sign-header Team-Api-Signature -headers Content-Type,Date
fi
exec "$@"
