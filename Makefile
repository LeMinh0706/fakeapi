# Binary names
API_BINARY=api

# Build directory
BUILD_DIR=build

# Docker image names
API_IMAGE=template-api

# Go build flags
BUILD_FLAGS=-v

# Postgres address
POSTGRES_URL=postgres://root:kocanpass@localhost:5432/$(dbname)?sslmode=disable

.PHONY: api clean run-api build-docker-api run-docker-api goose-create module goose-up goose-down sqlc protoc

api:
	go build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(API_BINARY) ./cmd/server

build-docker-api:
	docker build -t $(API_IMAGE):latest -f Dockerfile .

run-docker-api:
	docker run --rm -p 8080:8080 $(API_IMAGE):latest

clean:
	@rm -rf $(BUILD_DIR)

run-api: api
	./$(BUILD_DIR)/$(API_BINARY)

goose-create: 
	goose -dir db/schemas create $(name) sql	

goose-up:
	goose -dir db/schemas postgres $(POSTGRES_URL) up

goose-down:
	goose -dir db/schemas postgres $(POSTGRES_URL) down

sqlc:
	sqlc generate 

protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/proto/$(name)/*.proto

module:

