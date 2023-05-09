FROM golang:1.20.4-alpine3.16 as build

WORKDIR /app

RUN touch .env                                                                                                     
RUN printenv > .env 

COPY . .

RUN go mod tidy

RUN go build -o app -v .

ENTRYPOINT [ "./app" ]