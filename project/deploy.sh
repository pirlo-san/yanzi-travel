#!/bin/bash

set -e -u -x
make clean
make all

if [ $# -gt 0 ]; then
	PORT=$1
else
	PORT=80
fi

echo "start server container at port: $PORT"
docker run -d -p $PORT:$PORT --name server server:latest /root/gopath/bin/server -p $PORT
