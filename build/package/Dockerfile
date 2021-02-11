FROM golang:1.14

WORKDIR /app

COPY ../../deps.sh ./

RUN ./deps.sh

COPY ../../go.mod go.sum ./

RUN go mod download

COPY build/package .

RUN make build-linux

EXPOSE 9000

CMD ["./tmp/authenticator-server-linux"]
