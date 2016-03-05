#! /bin/bash -e

cd /usr/local/18f/team-api

LOGS=/var/log/team-api

if [ ! -d $LOGS ]; then
  mkdir -p $LOGS
fi

APP_SYS_ROOT=/usr/local/18f

source /etc/profile.d/nvm.sh

# For some reason, sourcing rbenv triggers an exit condition.
set +e
source /etc/profile.d/rbenv.sh
set -e

if [ ! -d team-api.18f.gov/.git ]; then
  git clone git@github.com:18F/team-api.18f.gov.git team-api.18f.gov
fi

cd team-api.18f.gov

# See https://github.com/18F/go_script/issues/16 about this business.
bundle install --path vendor/bundle
export GEM_HOME=$PWD/vendor/bundle/ruby/2.3.0
gem install bundler

git fetch origin master
git clean -f
git reset --hard origin/master
./go build
cd ..

forever start -l $LOGS/access.log -a \
  $NVM_BIN/team-api $APP_SYS_ROOT/team-api/config/team-api-config.json
