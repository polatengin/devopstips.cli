#!/usr/bin/env sh

# stop on error
# set -e

if ! type "docker" > /dev/null; then
  echo "Docker is not installed. Please install docker and try again."

  exit 1
fi
