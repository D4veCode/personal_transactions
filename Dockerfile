FROM golang:1.18-alpine as builder

WORKDIR /personal_transactions

COPY go.* ./
RUN go mod download

COPY . .

RUN go build .

EXPOSE 8080

ENTRYPOINT ["./personal_transactions"]
