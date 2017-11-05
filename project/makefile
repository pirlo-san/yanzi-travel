
.PHONY:  all clean server server_clean
.IGNORE:     clean        server_clean

all:   server
clean: server_clean

server:
	cd ../server
	docker build -t server:latest ../server
	cd $PWD

server_clean:
	docker kill $(docker ps -a -q)
	docker rm   $(docker ps -a -q)
	
