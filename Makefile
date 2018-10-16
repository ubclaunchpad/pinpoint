VERSION=`git rev-parse --short HEAD`
DEV_ENV=export `less ./dev/.env | xargs`
DEV_COMPOSE=docker-compose -f dev/docker-compose.yml

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
	go test -race -cover ./...
	( cd frontend ; npm run test -- --coverage )
	( cd client ; npm run test )

# Set up test environment
.PHONY: testenv
testenv:
	$(DEV_COMPOSE) up -d

# Stop test environment
.PHONY: testenv-stop
testenv-stop:
	$(DEV_COMPOSE) stop

# Clean up stuff
.PHONY: clean
clean: testenv-stop
	$(DEV_COMPOSE) rm -f -s -v

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
	protoc -I protobuf pinpoint.proto --go_out=plugins=grpc:protobuf
	make proto-pkg PKG=request
	make proto-pkg PKG=response

.PHONY: proto-pkg
proto-pkg:
	protoc -I protobuf $(PKG)/$(PKG).proto --go_out=plugins=grpc:$(GOPATH)/src

# Runs core service
.PHONY: core
core:
	go run core/main.go run --dev \
		--tls.cert dev/certs/127.0.0.1.crt \
		--tls.key dev/certs/127.0.0.1.key

# Runs API gateway
.PHONY: gateway
gateway:
	go run gateway/main.go run --dev \
		--core.cert dev/certs/127.0.0.1.crt

# Builds binary for pinpoint-core
.PHONY: pinpoint-core
pinpoint-core:
	go build -o ./bin/pinpoint-core \
    -ldflags "-X main.Version=$(VERSION)" \
    ./core

# Builds binary for pinpoint-gateway
.PHONY: pinpoint-gateway
pinpoint-gateway:
	go build -o ./bin/pinpoint-gateway \
    -ldflags "-X main.Version=$(VERSION)" \
    ./gateway
