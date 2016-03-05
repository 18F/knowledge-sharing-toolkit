#! /bin/bash -e

cd /usr/local/18f/lunr-server

LOGS=/var/log/lunr-server

if [ ! -d $LOGS ]; then
  mkdir -p $LOGS
fi

APP_SYS_ROOT=/usr/local/18f

source /etc/profile.d/nvm.sh

forever start -l $LOGS/access.log -a \
  $NVM_BIN/lunr-server $APP_SYS_ROOT/lunr-server/config/lunr-server-config.json
