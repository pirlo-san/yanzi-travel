
.PHONY:  all clean server server_clean
.IGNORE:     clean        server_clean

all:   server
clean: server_clean

server:
	tar -zcvf ../server/src.tar.gz -C ../server src
	docker build -t server:latest ../server

server_clean:
	rm -rf ../server/src.tar.gz
	docker kill `docker ps -a -q`
	docker rm   `docker ps -a -q`
	
