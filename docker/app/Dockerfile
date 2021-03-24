FROM golang:1.15-alpine as builder

WORKDIR /go/src/github.com/sk62793/todo_server

COPY . .
RUN go mod download
RUN go build .
ENV PATH /go/bin:$PATH

FROM golang:1.15-alpine
ENV DOCKERIZE_VERSION v0.6.1
WORKDIR /go/src/github.com/sk62793/todo_server

COPY --from=builder /go/src/github.com/sk62793/todo_server .
RUN chmod a+x todo_server

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz
