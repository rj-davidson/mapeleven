FROM golang:latest
WORKDIR /app
COPY ./api/.env /app/.env
COPY ./api /app/api
RUN mkdir -p /app/api/bin && \
    cd /app/api/cmd/db && \
    go build -o /app/api/bin/db ./main.go
ENTRYPOINT ["/app/api/bin/db", "-init"]
