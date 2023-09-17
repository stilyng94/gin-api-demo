#!/usr/bin/env sh
set -e

echo "start app"
./api
exec "$@"
