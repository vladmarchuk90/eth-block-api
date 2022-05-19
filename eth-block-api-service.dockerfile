# golang image
FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o ethBlockInfoApp ./cmd/web

RUN chmod +x /app/ethBlockInfoApp

CMD [ "/app/ethBlockInfoApp" ]
