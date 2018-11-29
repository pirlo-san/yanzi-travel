# This dockerfile builds the golang development environment,
# and then start our http server

FROM centos:7

MAINTAINER pirlo san <pirlo_san@163.com>

# install gcc
# -y means saying yes to all questions
RUN yum install -y gcc

# install golang
RUN yum install -y go

# config GOROOT
ENV GOROOT /usr/lib/golang
ENV PATH=$PATH:/usr/lib/golang/bin

# config GOPATH
RUN mkdir -p /root/gopath
RUN mkdir -p /root/gopath/src
RUN mkdir -p /root/gopath/pkg
RUN mkdir -p /root/gopath/bin
ENV GOPATH /root/gopath

# copy source files
RUN mkdir -p /root/gopath/src/server
COPY src/* /root/gopath/src/server/

# build the server
WORKDIR /root/gopath/src/server
RUN go build -o server.bin main.go

# startup the server
CMD /root/gopath/src/server/server.bin
