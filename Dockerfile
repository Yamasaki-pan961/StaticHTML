FROM golang:latest

WORKDIR /heroku

ENV GO111MODULE=on

COPY . /heroku

RUN go mod download

CMD ["go","run","server.go"]