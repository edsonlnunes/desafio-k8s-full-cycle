FROM golang:1.12-alpine AS builder

# RUN apk add --update gcc go && apk add bash

WORKDIR /go/src/

COPY . .

# RUN GOOS=linux go build -o main
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# CMD ["./main"]

FROM alpine:latest

WORKDIR /go/src/

COPY --from=builder /go/src/ .

CMD ["./main"]