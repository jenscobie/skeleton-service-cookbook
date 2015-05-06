#!/bin/bash

set -e -u

chef -v >/dev/null 2>&1 || { echo >&2 "Chef Development Kit is required. Please install the latest version."; exit 1; }

if [ $# -eq 0 ]
then
    ./gradlew tasks
else
    ./gradlew "$@"
fi
