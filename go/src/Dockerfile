FROM golang

ADD . /go/src/app

RUN go get golang.org/x/net/html

RUN go install app

ENTRYPOINT /go/bin/app

EXPOSE 8080