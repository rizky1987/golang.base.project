FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /

COPY . .

RUN go mod tidy

RUN go build main.go

ENTRYPOINT ["/main"]