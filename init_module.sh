#!/bin/sh

DIRNAME=$1

if [ -d "$DIRNAME" ]; then
  echo "Module $DIRNAME found. Delete the module before re-initializing"
  exit 1
else
  cp -r template $DIRNAME
fi
