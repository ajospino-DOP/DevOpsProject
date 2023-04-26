FROM golang:1.20

RUN apt-get update

WORKDIR /app

COPY . .

COPY .env .env

RUN go mod tidy

RUN go build -o app -v .

ENTRYPOINT [ "./app" ]