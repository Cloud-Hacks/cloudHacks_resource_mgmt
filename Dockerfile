## Build

FROM golang:1.16-alpine

WORKDIR /res-mgmt

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /web4api

## Deploy

EXPOSE 8080

CMD [ "/web4api" ]
