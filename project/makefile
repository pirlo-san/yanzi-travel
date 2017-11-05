
.PHONY: all clean server server_clean
PWD=`pwd`

all:   server
clean: server_clean

server:
	cd ../server
	docker build -t server:latest .
	cd $PWD

server_clean:
	docker kill $(docker ps -a -q)
	docker rm   $(docker ps -a -q)
	
