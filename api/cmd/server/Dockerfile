FROM golang:latest
WORKDIR /app/api
COPY ./api/go.mod ./api/go.sum ./
RUN go mod download
COPY ./api .
RUN go generate ./pkg/ent
RUN go build -o /app/api/bin/server ./cmd/server
WORKDIR /app/api/bin
COPY ./api/.env /app/api/bin
EXPOSE 8080
ENTRYPOINT ["./server"]
