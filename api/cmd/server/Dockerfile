# syntax=docker/dockerfile:1
FROM golang:latest
WORKDIR /app/
COPY api/cmd/server/go.mod api/cmd/server/go.sum ./
RUN go mod download
COPY . .
RUN go generate ./pkg/ent
RUN go build -o /app/server ./api/cmd/server
EXPOSE 8080
ENTRYPOINT ["/app/server"]