#!/bin/bash

BASE=$(dirname $(readlink -f $0))
echo $BASE
REPO=$1

# using https://github.com/baxterthehacker/public-repo as a test repository

cd $REPO \
    && git checkout origin/master \
    && git branch -f master 7b80eb100206a56523dbda6202d8e5daa05e265b \
    && git checkout master \
    && curl --data "@$BASE/core/testdata/pushevent.json" http://127.0.0.1:8000/deploy
