FROM golang:1.10.2

RUN mkdir -p /go/src/github.com/stack360/faas-mq/

WORKDIR /go/src/github.com/stack360/faas-mq

COPY vendor      vendor
COPY handlers    handlers
COPY mq          mq
COPY server.go   .

RUN gofmt -l -d $(find . -type f -name '*.go' -not -path "./vendor/*") \
  && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o faas-mq .

FROM alpine:3.5
RUN apk --no-cache add ca-certificates
WORKDIR /root/

EXPOSE 8080
ENV http_proxy      ""
ENV https_proxy     ""

COPY --from=0 /go/src/github.com/stack360/faas-mq/faas-mq    .

CMD ["./faas-mq"]
