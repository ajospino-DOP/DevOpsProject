FROM golang:1.20

RUN apt-get update

ARG MONGODB_URI
ENV MONGODB_URI = ${MONGODB_URI}

RUN touch /.env                                                                                                     
RUN printenv > /.env 

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o app -v .

ENTRYPOINT [ "./app" ]