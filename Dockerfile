FROM artifactory.accuknox.com/accuknox/golang:1.17-alpine3.15 AS builder
WORKDIR /home/resource-service
COPY . .
RUN go mod tidy
RUN GOOS=linux go build -o main .

FROM artifactory.accuknox.com/accuknox/alpine:3.15
WORKDIR /home/resource-service
# ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
# ENV ZONEINFO /zoneinfo.zip
# RUN apt-get update && apt-get install -y apt-transport-https ca-certificates
COPY --from=builder /home/resource-service/main /home/resource-service/main
CMD ["./main"]
