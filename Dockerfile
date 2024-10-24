FROM golang:1.23-alpine AS builder

WORKDIR /go/app/sports

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go/app/sports /main

ENV APP_ENV=production

CMD ["/main"]