#! /bin/bash

cd /usr/local/18f

echo "Starting the system..."

echo "Sync 18F pages sites:"
pages/run.sh sync-data

echo "Runing daemon services:"
oauth2_proxy/run-server.sh
hmacproxy/run.sh run-server
authdelegate/run-server.sh
pages/run.sh run-server
lunr-server/run-server.sh
team-api/run-server.sh
nginx/run-server.sh

echo "System start complete."
