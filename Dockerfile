FROM golang:1.20

RUN apt-get update

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy

EXPOSE 8060

RUN go main.go