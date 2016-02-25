#! /bin/bash -e

source /etc/profile.d/gvm.sh

if [ "$1" = "run-server" ]; then
  exec authdelegate config/authdelegate-config.json
fi
exec "$@"
