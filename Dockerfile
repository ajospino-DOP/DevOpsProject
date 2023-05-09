FROM golang:1.20.4 as build

WORKDIR /app

RUN touch .env                                                                                                     
RUN printenv > .env 

COPY . .

RUN go mod tidy

RUN go build -o app -v .

#Multistage to avoid image size being unmanageable - try 2

FROM alpine:latest as main

COPY --from=build /app .
COPY --from=build /go/bin/ /bin/

RUN ls -l

ENTRYPOINT [ "./app" ]