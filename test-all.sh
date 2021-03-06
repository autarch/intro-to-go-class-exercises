#!/bin/bash

set -e

for dir in */answer/*; do
    if [[ ! -d $dir ]]; then
        continue
    fi
    pushd $dir
    if [[ -n $(grep 'package main' *.go) ]]; then
        go build
    fi
    if [[ $dir =~ "6-unit-tests" ]]; then
        set +e
    fi
    go test -v
    set -e
    popd
done
