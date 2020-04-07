FROM golang:1.12-alpine

# RUN apk add --update gcc go && apk add bash

WORKDIR /go/src/

COPY ./src .

RUN GOOS=linux go build -o main

CMD [ "./main" ]
