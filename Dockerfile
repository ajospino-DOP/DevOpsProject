FROM golang:1.20

RUN apt-get update

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o app -v ./..

ENTRYPOINT [ "./app" ]