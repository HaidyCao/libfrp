#!/bin/bash

unset GO111MODULE
unset GOMOD

if [ ! -d "${ANDROID_HOME}/ndk-bundle" ] && [ ! -d "${ANDROID_NDK_HOME}" ]; then
  if [ "Darwin" == "$(uname)" ]; then
    ANDROID_NDK_HOME="${HOME}/Library/Android/sdk/ndk/21.3.6528147"
  else
    exit
  fi
fi

go env -w GO111MODULE=auto
gomobile bind -v -o libfrp.aar -target=android .
