FROM golang:1.16-alpine3.15 AS builder
WORKDIR /home/resource-mgmt
COPY . .
RUN go mod tidy
RUN GOOS=linux go build -o main .

FROM alpine:3.15
WORKDIR /home/resource-mgmt

COPY --from=builder /home/resource-mgmt/main /home/resource-mgmt/main
CMD ["./main"]
