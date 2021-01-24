FROM golang:1.15.2-alpine

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN apk add --no-cache alpine-sdk git \
    && go get github.com/pilu/fresh \
    && go get bitbucket.org/liamstask/goose/cmd/goose

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

CMD ["fresh"]