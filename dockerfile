FROM golang:1.14.0

WORKDIR /app

COPY ./ /app

RUN go mod download

ENV POSTGRES_USER=${POSTGRES_USER}
ENV POSTGRES_PASSWORD=${POSTGRES_PASSWORD}
ENV POSTGRES_DB=${POSTGRES_DB}
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}

RUN go build -o go-demo cmd/web/main.go

CMD ./go-demo -migrate