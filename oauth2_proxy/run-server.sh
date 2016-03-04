#! /bin/bash -e

cd /usr/local/18f/oauth2_proxy

LOGS=/var/log/oauth2_proxy

if [ ! -d $LOGS ]; then
  mkdir -p $LOGS
fi

. /etc/profile.d/gvm.sh
. /usr/local/18f/oauth2_proxy/config/env-secret.sh

nohup $GOPATH/bin/oauth2_proxy \
  --config=/usr/local/18f/oauth2_proxy/config/oauth2_proxy.cfg \
  >>$LOGS/access.log 2>&1 &
