#! /bin/bash -e

LOGS=/var/log/nginx

if [ ! -d $LOGS ]; then
  mkdir -p $LOGS
fi

sudo /usr/local/18f/nginx/sbin/nginx
