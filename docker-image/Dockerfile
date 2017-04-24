FROM golang
MAINTAINER Edgar Rios (edgar.gs@gmail.com)
COPY Hello.go /go/src/github.com/edgargs/MyHelloWorldGolang/
RUN go install github.com/edgargs/MyHelloWorldGolang
ENV HELLO_GO_NAME Edgar
ENTRYPOINT /go/bin/MyHelloWorldGolang

