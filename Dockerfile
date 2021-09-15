FROM golang:1.17.1

RUN mkdir /app
WORKDIR /app

COPY . .
RUN go mod download
