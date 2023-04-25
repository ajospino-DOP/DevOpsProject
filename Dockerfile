FROM golang:1.20

WORKDIR /usr/src/app

COPY . /usr/src/app/


RUN go mod tidy







RUN go main.go