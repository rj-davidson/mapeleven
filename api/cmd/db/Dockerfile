# Use an official Go runtime as a parent image
FROM golang:latest
WORKDIR /app
COPY ./bin/db /app/db
EXPOSE 80
CMD ["./db -initialize"]
