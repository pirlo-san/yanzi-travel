#!/bin/bash

set -e -u -x
make clean
make all
docker run -d -p $1:$1 --name server server:latest
