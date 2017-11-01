# This dockerfile build the golang development environment,
# and then start our http server

FROM centos:7

MAINTAINER pirlo san <pirlo_san@163.com>

# install gcc
# -y means saying yes to all questions
RUN yum install -y gcc

# install golang
RUN rm -rf /usr/local/go
RUN mkdir -p /usr/local/go
RUN yum install --installroot=/usr/local/go --releasever=/ -y go

# config GOROOT
ENV GOROOT /usr/local/go
ENV PATH=$PATH:/usr/local/go/bin

# config GOPATH
RUN mkdir -p /root/gopath
RUN mkdir -p /root/gopath/src
RUN mkdir -p /root/gopath/pkg
RUN mkdir -p /root/gopath/bin
ENV GOPATH /root/gopath

# copy source files
COPY src /root/gopath/server

# startup the server
RUN cd /root/gopath/server
RUN /usr/local/go/bin/go run main.go
