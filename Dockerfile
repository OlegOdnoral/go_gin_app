FROM golang:latest

LABEL app_name="golang gin test app"

LABEL maintainer="Oleg_Odnoral"

ADD . /go/src/github.com/OlegOdnoral/go_gin_app

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/OlegOdnoral/go_gin_app

RUN dep ensure

RUN go build ./src/main.go

ENTRYPOINT ./main

EXPOSE 8080