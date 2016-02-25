#! /bin/bash -e

source /etc/profile.d/nvm.sh
source /etc/profile.d/pyenv.sh

# For some reason, sourcing rbenv triggers an exit condition.
set +e
source /etc/profile.d/rbenv.sh
set -e

source config/env-secret.sh

if [ "$1" = "run-server" ]; then
  aws s3 sync s3://18f-pages/sites $APP_SYS_ROOT/pages/sites && \
      18f-pages config/pages-config.json
  exec 18f-pages config/pages-config.json
fi
exec "$@"
