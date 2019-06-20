FROM golang:1.8

WORKDIR /go/src/github.com/vavas/go-parser
COPY . /go/src/github.com/vavas/go-parser

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]