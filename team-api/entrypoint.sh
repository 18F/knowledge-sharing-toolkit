#! /bin/bash -e

if [ -f $APP_SYS_ROOT/ssh/config/prep-ssh.sh ]; then
  source $APP_SYS_ROOT/ssh/config/prep-ssh.sh
fi

source /etc/profile.d/nvm.sh

# For some reason, sourcing rbenv triggers an exit condition.
set +e
source /etc/profile.d/rbenv.sh
set -e

if [ "$1" = "run-server" ]; then
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

  exec team-api config/team-api-config.json
fi
exec "$@"
