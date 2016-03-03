#! /bin/bash -e

source /etc/profile.d/nvm.sh

# For some reason, sourcing rbenv triggers an exit condition.
set +e
source /etc/profile.d/pyenv.sh
source /etc/profile.d/rbenv.sh
set -e

if [ -f $APP_SYS_ROOT/ssh/config/prep-ssh.sh ]; then
  source $APP_SYS_ROOT/ssh/config/prep-ssh.sh
fi

if [ -f config/env-secret.sh ]; then
  source config/env-secret.sh
fi

function sync_data() {
  aws s3 sync s3://18f-pages/sites $APP_SYS_ROOT/pages/sites --delete
}

if [ "$1" = "sync-data" ]; then
  sync_data
  exit
fi

if [ "$1" = "run-server" ]; then
  sync_data
  exec 18f-pages config/pages-config.json
fi
exec "$@"
