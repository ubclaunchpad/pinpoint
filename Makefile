VERSION=`git rev-parse --short HEAD`
DEV_ENV=export `less ./dev/.env | xargs`
TEST_COMPOSE=docker-compose -f dev/testenv.yml -p test
MON_COMPOSE=docker-compose -f dev/monitoring.yml -p monitoring

all: check

# Run simple checks
.PHONY: check
check:
	go vet ./...
	go test -run xxxx ./...

# Install dependencies
.PHONY: deps
deps:
	go get -u github.com/maxbrunsfeld/counterfeiter
	go get -u github.com/vburenin/ifacemaker
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
	mkdir -p tmp/data
	$(TEST_COMPOSE) up -d

# Stop test environment
.PHONY: testenv-stop
testenv-stop:
	$(TEST_COMPOSE) stop

# Set up monitoring environment
.PHONY: monitoring
monitoring:
	mkdir -p tmp/data
	$(MON_COMPOSE) up -d

# Stop monitoring environment
.PHONY: monitoring-stop
monitoring-stop:
	$(MON_COMPOSE) stop

# Clean up stuff
.PHONY: clean
clean: testenv-stop monitoring-stop
	$(TEST_COMPOSE) rm -f -s -v
	$(MON_COMPOSE) rm -f -s -v
	rm -rf tmp

# Run linters and checks
.PHONY: lint
lint: SHELL:=bash
lint: check
	diff -u <(echo -n) <(gofmt -d -s `find . -type f -name '*.go' -not -path "./vendor/*"`)
	diff -u <(echo -n) <(golint `go list ./... | grep -v /vendor/`)
	( cd frontend ; npm run lint )
	( cd client ; npm run lint )

# Regenerate all generated code
.PHONY: gen
gen: proto mocks

# Generate protobuf code from definitions
.PHONY: proto
proto:
	protoc -I protobuf pinpoint.proto --go_out=plugins=grpc:protobuf
	make proto-pkg PKG=request
	make proto-pkg PKG=response
	# generate mock
	counterfeiter -o ./protobuf/fakes/pinpoint.pb.go \
		./protobuf/pinpoint.pb.go CoreClient

.PHONY: proto-pkg
proto-pkg:
	protoc -I protobuf $(PKG)/$(PKG).proto --go_out=plugins=grpc:$(GOPATH)/src

.PHONY: mocks
mocks:
	# generate database interface and mock 
	ifacemaker \
		-f ./core/database/*.go \
		-s Database \
		-i DBClient \
		--pkg database \
		-o ./core/database/database.i.go \
		-c "Code generated by ifacemaker. DO NOT EDIT." \
		-y "DBClient wraps the AWS DynamoDB database API"
	counterfeiter -o ./core/database/mocks/database.i.go \
		./core/database/database.i.go DBClient

# Runs core service
.PHONY: core
core:
	go run core/main.go run --dev \
		--tls.cert dev/certs/127.0.0.1.crt \
		--tls.key dev/certs/127.0.0.1.key $(FLAGS)

# Runs API gateway
.PHONY: gateway
gateway:
	go run gateway/main.go run --dev \
		--core.cert dev/certs/127.0.0.1.crt $(FLAGS)

.PHONY: gateway-tls
gateway-tls:
	go run gateway/main.go run --dev \
		--core.cert dev/certs/127.0.0.1.crt \
		--tls.cert dev/certs/127.0.0.1.crt \
		--tls.key dev/certs/127.0.0.1.key $(FLAGS)

# Runs web app
.PHONY: web
web:
	( cd frontend ; npm start )

# Builds binary for pinpoint-core
.PHONY: pinpoint-core
pinpoint-core:
	go build -o ./bin/pinpoint-core \
    -ldflags "-X main.Version=$(VERSION)" \
    ./core $(FLAGS)

# Builds binary for pinpoint-gateway
.PHONY: pinpoint-gateway
pinpoint-gateway:
	go build -o ./bin/pinpoint-gateway \
    -ldflags "-X main.Version=$(VERSION)" \
    ./gateway $(FLAGS)
