#!/bin/bash

set -e -u -x
make clean
make all
docker run -d -p 80:80 --name server server:latest
