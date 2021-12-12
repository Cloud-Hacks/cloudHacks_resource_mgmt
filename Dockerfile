FROM golang:1.16 as builder

WORKDIR /res-mgmt

ADD . /res-mgmt/

RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go



FROM alpine:3.7

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /root

COPY --from=builder /res-mgmt/. .

CMD ["./main"]
