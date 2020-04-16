#!/bin/sh

if [ ${MODE} == 'local' ]; then
  docker-compose -f deploy/docker-compose.local.yml up -d
elif [ ${MODE} == 'release' ]; then
  docker-compose -f deploy/docker-compose.yml up -d
fi
