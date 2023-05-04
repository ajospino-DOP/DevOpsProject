FROM golang:1.20 as build

RUN apt-get update

ARG MONGODB_URI
ENV MONGODB_URI = ${MONGODB_URI}

RUN touch .env                                                                                                     
RUN printenv > .env 

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o app -v .

#Multistage to avoid image size being unmanageable - try 2

FROM golang:1.20.4-alpine3.16 as main

COPY --from=build /app /

RUN ls -a /src

RUN echo 1

RUN ls -a /bin

ENTRYPOINT [ "app" ]