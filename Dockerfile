FROM golang:latest

LABEL author="zhensheng.five@gmail.com"

WORKDIR $GOPATH/src/github.com/everywan/go-web

ADD . $GOPATH/src/github.com/everywan/go-web
RUN make dist

EXPOSE 50000

ENTRYPOINT ["./bin/main"]
