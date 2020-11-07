#!/bin/bash

unset GO111MODULE
unset GOMOD

go env -w GO111MODULE=auto
gomobile bind -v -target=ios .
