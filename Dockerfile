FROM golang:alpine as build

WORKDIR /app

RUN touch .env                                                                                                     
RUN printenv > .env 

COPY . .

RUN go mod tidy

RUN go build -o app .

FROM alpine as main

COPY --from=build /app /

ENTRYPOINT [ "./app" ]
