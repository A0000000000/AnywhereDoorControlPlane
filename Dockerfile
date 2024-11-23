FROM ubuntu:latest
LABEL authors="maoyanluo"

WORKDIR /ws

ADD https://dl.google.com/go/go1.23.3.linux-amd64.tar.gz /ws

RUN tar -zxvf go1.23.3.linux-amd64.tar.gz

COPY src /ws/code

RUN mkdir GOPATH

ENV GOROOT=/ws/go
ENV GOPATH=/ws/GOPATH

ENV PATH=$PATH:${GOROOT}/bin

WORKDIR /ws/code

RUN go get

RUN go build -o main

EXPOSE 80

ENTRYPOINT ["./main"]