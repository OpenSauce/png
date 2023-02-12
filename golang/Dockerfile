# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD . .

RUN go build -o /golang cmd/main.go

EXPOSE 3000

CMD [ "/golang" ]