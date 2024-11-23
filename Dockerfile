FROM ubuntu:latest
LABEL authors="maoyanluo"

WORKDIR /ws

#ADD https://dl.google.com/go/go1.23.3.linux-amd64.tar.gz /ws
ADD https://mirrors.aliyun.com/golang/go1.23.3.linux-amd64.tar.gz /ws

RUN tar -zxvf go1.23.3.linux-amd64.tar.gz

COPY src /ws/code

RUN mkdir GOPATH

ENV GOROOT=/ws/go
ENV GOPATH=/ws/GOPATH

ENV PATH=$PATH:${GOROOT}/bin

WORKDIR /ws/code

# 解决CA证书问题
RUN apt update
RUN apt install -y --no-install-recommends ca-certificates

RUN go get

RUN go build -o main

EXPOSE 80

ENTRYPOINT ["./main"]