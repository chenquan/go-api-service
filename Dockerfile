FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/opensourceai/go-api-service
COPY . $GOPATH/src/github.com/opensourceai/go-api-service
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-api-service"]
