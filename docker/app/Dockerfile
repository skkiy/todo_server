FROM golang:1.8
RUN apt-get update -qq && \
  apt-get install -y mysql-client vim

WORKDIR /go/src/todo/todo_server

ADD . .
RUN go build
ENV PATH /go/bin:$PATH
