#!/usr/bin/env sh

# stop on error
# set -e

if ! type "docker" > /dev/null; then
  echo "Docker is not installed. Please install docker and try again."

  exit 1
fi

if ! type "devopstips" > /dev/null; then
  echo "Previous devopstips installation doesn't found, installing fresh..."

  touch ~/devopstips
  chmod +x ~/devopstips

  "alias devopstips=\"~/devopstips\"" >> ~/.bashrc

  source ~/.bashrc
fi

"docker run --rm -it devopstips:0.1.0" > ~/devopstips

