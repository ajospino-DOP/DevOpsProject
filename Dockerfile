FROM golang:1.20.4-alpine3.16 as build

WORKDIR /app

RUN touch .env                                                                                                     
RUN printenv > .env 

COPY . .

RUN go mod tidy

RUN go build -o app -v .

FROM alpine:3.16 as main

COPY --from=build /app /

ENTRYPOINT [ "./app" ]

#re testing dockerfile