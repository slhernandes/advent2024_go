#!/bin/sh

DIRNAME=$1

if [ -d "${DIRNAME}" ]; then
  echo -n "You are about to delete ${DIRNAME}. Are you sure? [y/N] "
  read input
  case $input in
    y|Y)
      rm -r ${DIRNAME}
      ;;
    *)
      echo "Canceled deinit for ${DIRNAME}."
      exit 1
      ;;
  esac
else
  echo "Directory ${DIRNAME} not found."
fi
