FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download

ENV POSTGRES_USER=${POSTGRES_USER}
ENV POSTGRES_PASSWORD=${POSTGRES_PASSWORD}
ENV POSTGRES_DB=${POSTGRES_DB}

RUN go build -o go-demo cmd/web/main.go

ENTRYPOINT ./go-demo
