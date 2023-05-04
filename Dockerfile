FROM golang:1.20 as build

RUN apt-get update

ARG MONGODB_URI
ENV MONGODB_URI = ${MONGODB_URI}

RUN touch .env                                                                                                     
RUN printenv > .env 

WORKDIR /app

COPY . .

RUN go mod tidy

FROM golang:1.20.4-alpine3.16 as main

COPY --from=build /app /

RUN go build -o app -v .

ENTRYPOINT [ "./app" ]