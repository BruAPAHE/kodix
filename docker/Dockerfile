FROM golang:latest

RUN set -xe \
    && apt-get update && apt-get install -y \
       build-essential \
    && apt-get clean  \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

COPY ./src/ /go/src/app/

WORKDIR /go/src/app/

RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/geography/main.go \
    && cp app $PATH_BIN/app \
    && chmod 755 $PATH_BIN/app

ENTRYPOINT ["/usr/local/bin/app"]

EXPOSE 80