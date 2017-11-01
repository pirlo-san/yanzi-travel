#!/bin/bash

PWD=`pwd`
DATABASE_DIR=../mysql

function make_database {
	cd $DATABASE_DIR
	chmod 777 make.sh
	./make.sh
	cd $PWD
}

set -x
make_database;
