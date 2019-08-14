FROM golang:1.12.8

WORKDIR /app

COPY deps.sh ./

RUN ./deps.sh

COPY go.mod go.sum ./

RUN go mod download

EXPOSE 9000
