FROM golang:alpine AS build

WORKDIR /app

RUN touch .env                                                                                                     
RUN printenv > .env 

COPY . .

RUN go mod tidy

RUN go build -o app .

FROM alpine AS main

COPY --from=build /app /

ENTRYPOINT [ "./app" ]
