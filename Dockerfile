FROM golang:1.22.5-alpine3.20 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . ./
RUN go build cmd/api/main.go
CMD [ "./main" ]

