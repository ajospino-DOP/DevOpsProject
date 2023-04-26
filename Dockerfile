FROM golang:1.20

RUN apt-get update

ARG MONGODB_URI

WORKDIR /app

COPY . .

RUN touch .env
RUN echo ${MONGODB_URI} > .env

RUN go mod tidy

RUN go build -o app -v .

ENTRYPOINT [ "./app" ]