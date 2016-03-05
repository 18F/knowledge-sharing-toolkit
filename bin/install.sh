#! /bin/sh

/usr/local/18f/bin/dev-base.sh
/usr/local/18f/bin/dev-standard.sh

/usr/local/18f/oauth2_proxy/install.sh
/usr/local/18f/hmacproxy/install.sh
/usr/local/18f/authdelegate/install.sh
/usr/local/18f/pages/install.sh
/usr/local/18f/lunr-server/install.sh
/usr/local/18f/team-api/install.sh
/usr/local/18f/nginx/install.sh

for service in \
  oauth2_proxy \
  hmacproxy \
  authdelegate \
  pages \
  lunr-server \
  team-api \
  nginx
do
  logdir=/var/log/$service
  if [ ! -d $logdir ]; then
    sudo mkdir $logdir
    sudo chown ubuntu:ubuntu $logdir
  fi
done
