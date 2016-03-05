#! /bin/bash -e

cd /usr/local/18f/authdelegate

LOGS=/var/log/authdelegate

if [ ! -d $LOGS ]; then
  mkdir -p $LOGS
fi

source /etc/profile.d/gvm.sh

nohup $GOPATH/bin/authdelegate config/authdelegate-config.json \
  >>$LOGS/access.log 2>&1 &
