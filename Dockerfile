FROM golang:1.8

WORKDIR /GolandProjects/go-parser
COPY . /GolandProjects/go-parser

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]