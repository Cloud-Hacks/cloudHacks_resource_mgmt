FROM golang:1.16 as builder

WORKDIR /resource-mgmt-api

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 go build -o main .

FROM alpine:3.15

WORKDIR /resource-mgmt-api

COPY --from=builder /resource-mgmt-api/. .

CMD ["./main"]