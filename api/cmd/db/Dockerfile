FROM golang:latest
WORKDIR /app
COPY ./api /app/api
RUN mkdir -p /app/api/bin && \
    cd /app/api/cmd/db && \
    go build -o /app/api/bin/db
ENTRYPOINT ["/app/api/bin/db", "-initialize"]
