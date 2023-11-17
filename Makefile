GOBIN := $(shell go env GOPATH)/bin
ENV_FILE := ./api/.env
include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE))

all: generate build

generate:
	@echo "+ $@"
	cd api && go generate ./...

build: db server client

db:
	@echo "+ $@"
	cd api/cmd/db && go build -o ../../bin/db ./fetcher.go

server:
	@echo "+ $@"
	cd api/cmd/server && go build -o ../../bin/server

client:
	@echo "+ $@"
	cd client && yarn install && yarn build

docker-build: docker-build-api-db docker-build-api-server docker-build-client

docker-build-api: db docker-build-api-db docker-build-api-server

docker-up-api: docker-build-api
	@echo "+ $@"
	docker-compose --env-file $(ENV_FILE) up db api-db api-server

docker-up-api-detached: docker-build-api
	@echo "+ $@"
	docker-compose --env-file $(ENV_FILE) up -d db api-db api-server

docker-build-api-db: db
	@echo "+ $@"
	docker-compose build api-db

docker-build-api-server:
	@echo "+ $@"
	docker-compose build api-server

docker-build-client:
	@echo "+ $@"
	docker-compose build client

docker-up: docker-build
	@echo "+ $@"
	docker-compose --env-file $(ENV_FILE) up

docker-up-db-detached:
	@echo "+ $@"
	docker-compose --env-file $(ENV_FILE) up db -d

docker-down:
	@echo "+ $@"
	docker-compose down

docker-down-db:
	@echo "+ $@"
	docker-compose down db

docker-up-detached: docker-build
	@echo "+ $@"
	docker-compose up -d

clean:
	@echo "+ $@"
	rm -f api/bin/db api/bin/server
	cd client && rm -rf build

.PHONY: all generate build db server client docker-build docker-build-api-db docker-build-api-server docker-build-client docker-up docker-down clean
