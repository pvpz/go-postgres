FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download

EXPOSE 8080
ENTRYPOINT go run main.go