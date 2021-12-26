FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/hash_url
COPY . .

RUN go mod tidy
RUN go build -o /go/bin/hash_url .

FROM alpine
COPY --from=builder /go/bin/hash_url /usr/bin/hash_url

EXPOSE 1103

ENTRYPOINT ["/usr/bin/hash_url"]