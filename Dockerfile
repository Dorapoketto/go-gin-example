FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/Dorapoketto/go-gin-example
COPY . $GOPATH/src/github.com/Dorapoketto/go-gin-example
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./go-gin-example"]