FROM golang:latest

LABEL author="zhensheng.five@gmail.com"

WORKDIR $GOPATH/src/github.com/everywan/go-web-demo

ADD . $GOPATH/src/github.com/everywan/go-web-demo
RUN make dist

EXPOSE 50000

ENTRYPOINT ["./bin/main"]
