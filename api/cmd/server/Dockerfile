FROM golang:latest AS builder
WORKDIR /app
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
COPY ./api/go.mod ./api/go.sum ./
RUN go mod download
COPY ./api .
RUN go generate ./pkg/ent
RUN go build -o /app/server ./cmd/server

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
ENTRYPOINT ["./server"]
