FROM golang:1.12-alpine

WORKDIR /go/src/

COPY . .

RUN GOOS=linux go build -o main

CMD [ "./main" ]
