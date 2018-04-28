FROM golang:1.10.1-alpine3.7 

ADD . $GOPATH/src/futureHealth

WORKDIR $GOPATH/src/futureHealth

RUN go build -o futureHealth . 

EXPOSE 8080

CMD ["./futureHealth"]