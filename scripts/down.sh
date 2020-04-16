#!/bin/sh

if [ ${MODE} == 'local' ]; then
  docker-compose -f deploy/docker-compose.local.yml down
elif [ ${MODE} == 'release' ]; then
  docker-compose -f deploy/docker-compose.yml down
fi
