
.PHONY:  all clean server server_clean
.IGNORE:     clean        server_clean

all:   server
clean: server_clean

server:
	docker build -t server:latest ../server

server_clean:
	docker kill $(docker ps -a -q)
	docker rm   $(docker ps -a -q)
	
