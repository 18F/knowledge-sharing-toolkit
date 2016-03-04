#! /bin/bash -e

cd /usr/local/18f/pages

LOGS=/var/log/pages

if [ ! -d $LOGS ]; then
  mkdir -p $LOGS
fi

APP_SYS_ROOT=/usr/local/18f
source /etc/profile.d/nvm.sh

# For some reason, sourcing rbenv triggers an exit condition.
set +e
source /etc/profile.d/pyenv.sh
source /etc/profile.d/rbenv.sh
set -e

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
  forever start -l $LOGS/pages.log -a \
    $NVM_BIN/18f-pages $APP_SYS_ROOT/pages/config/pages-config.json
fi
