#!/bin/bash

MODULES="server"

function make_module {
	PWD=`pwd`
	cd ../$1
	chmod 777 make.sh
	./make.sh
	cd $PWD
}

function make_clean {
	docker kill $(docker ps -a -q)
	docker rm   $(docker ps -a -q)
}

if [ $# != 1 ] ; then
	echo "usage: $0 <module_name/clean/all>"
	exit 1
fi 

if [ $1 = "all" ] ; then
	for module in $MODULES  
	do   
		make_module $module
	done  
elif [ $1 = "clean" ] ; then
	make_clean
else
	make_module $1
fi
