VERSION=`git rev-parse --short HEAD`
DEV_ENV=export `less ./dev/.env | xargs`

all: check

# Run simple checks
.PHONY: check
check:
	go vet ./...
	go test -run xxxx ./...

# Install dependencies
.PHONY: deps
deps:
	dep ensure
	( cd frontend ; npm install )
	( cd client ; npm install )

# Execute tests
.PHONY: test
test:
	go test -race ./...
	( cd frontend ; npm run test -- --coverage )
	( cd client ; npm run test )

# Set up test environment
testenv:
	docker-compose -f dev/docker-compose.yml up -d

# Run linters and checks
.PHONY: lint
lint: check
	go fmt ./...
	golint `go list ./... | grep -v /vendor/`
	( cd frontend ; npm run lint )
	( cd client ; npm run lint )

# Generate protobuf code from definitions
.PHONY: proto
proto:
	protoc -I protobuf pinpoint.proto --go_out=plugins=grpc:grpc
	make proto-pkg PKG=request
	make proto-pkg PKG=response

.PHONY: proto-pkg
proto-pkg:
	protoc -I protobuf $(PKG).proto --go_out=plugins=grpc:$(GOPATH)/src

# Runs core service
.PHONY: core
core:
	go build -o ./bin/pinpoint-core \
    -ldflags "-w -s -X main.Version=$(VERSION)" \
    ./core

# Runs gateway api server
.PHONY: gateway
gateway:
	go build -o ./bin/pinpoint-gateway \
    -ldflags "-w -s -X main.Version=$(VERSION)" \
    ./gateway
