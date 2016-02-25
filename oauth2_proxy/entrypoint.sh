#! /bin/bash -e

source /etc/profile.d/gvm.sh
source config/env-secret.sh

if [ "$1" = "run-server" ]; then
  exec oauth2_proxy --config=config/oauth2_proxy.cfg
fi
exec "$@"
