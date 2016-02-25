#! /bin/bash -e

source /etc/profile.d/nvm.sh

if [ "$1" = "run-server" ]; then
  exec lunr-server config/lunr-server-config.json
fi
exec "$@"
