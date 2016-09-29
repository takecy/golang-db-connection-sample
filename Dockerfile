FROM golang:1.7.1-alpine

MAINTAINER takecy

COPY . $GOPATH/src/github.com/takecy/golang-mongo-sample

WORKDIR $GOPATH/src/github.com/takecy/golang-mongo-sample

CMD ["go", "run", "main.go"]
