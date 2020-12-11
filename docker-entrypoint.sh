#!/bin/bash
set -e

if [ "$1" = 'grpc_mock' ]; then
  if [ ! -f "$(go env GOPATH)/bin/grpc_mock" ]; then
    bash build.sh
  fi
fi

exec "$@"
