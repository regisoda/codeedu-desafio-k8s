FROM golang:1.10-alpine3.8 AS multistage

RUN apk add --no-cache --update git

WORKDIR /go/src/app

COPY . .

RUN go get -d -v \
    && go install -v \
    && go build

FROM alpine:3.8

COPY --from=multistage /go/bin/app /go/bin/

EXPOSE 8080

CMD ["/go/bin/app"]
