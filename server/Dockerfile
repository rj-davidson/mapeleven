# syntax=docker/dockerfile:1
FROM golang:latest
WORKDIR /app/
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go generate ./models/ent && go build -o main .
EXPOSE 8080
CMD ["/app/main"]
