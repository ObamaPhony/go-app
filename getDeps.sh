#!/bin/bash

DEPS=("github.com/inconshreveable/log15"
      "github.com/jeffail/gabs"
      "github.com/julienschmidt/httprouter"
      "gopkg.in/mgo.v2"
      "gopkg.in/pipe.v2")

# Get each dependency.

for dep in "${DEPS[@]}"; do
    go get -v -u "${dep}"
done
