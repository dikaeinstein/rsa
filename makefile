## Run tests with coverage
test-cover:
	GO111MODULE=on go test -coverprofile=cover.out -race -v ./...

## Run tests
test:
	GO111MODULE=on go test -race -v ./...

## Build binary
build:
	GO111MODULE=on go build *.go
