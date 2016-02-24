#! /bin/bash -e

if [ ! -d team-api.18f.gov/.git ]; then
  git clone git@github.com:18F/team-api.18f.gov.git team-api.18f.gov
fi

cd team-api.18f.gov
git fetch origin master
git clean -f
git reset --hard origin/master
./go build
