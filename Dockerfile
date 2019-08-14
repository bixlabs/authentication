FROM golang

WORKDIR /app

COPY deps.sh ./

RUN ./deps.sh

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make build-for-docker

EXPOSE 9000
