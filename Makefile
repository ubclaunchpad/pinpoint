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

# Execute tests
test:
	go test ./...

# Generate protobuf code from definitions
.PHONY: proto
proto:
	protoc -I protobuf pinpoint.proto --go_out=plugins=grpc:grpc
	make proto-pkg PKG=request
	make proto-pkg PKG=response

.PHONY: proto-pkg
proto-pkg:
	protoc -I protobuf $(PKG).proto --go_out=plugins=grpc:$(GOPATH)/src
