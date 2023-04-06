# syntax=docker/dockerfile:1
FROM golang:latest
WORKDIR /app/server/
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o main .
EXPOSE 8080
CMD ["/app/server/main"]