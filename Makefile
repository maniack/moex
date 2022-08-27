BINARY=trade
VERSION=$(shell git describe --tags --always)
BUILD=$(shell git describe --tags --long --always)
LDFLAGS=-ldflags "-X main.Version=$(BUILD)"
LINUX=$(BINARY)_linux_amd64
DARWIN=$(BINARY)_darwin_amd64
WINDOWS=$(BINARY)_windows_amd64.exe

default: all

all: get test linux darwin windows

get:
	go get
	go mod tidy
	go mod vendor

test:
	go test ./...

linux:
	env GOOS=linux GOARCH=amd64 go build -v -o dist/$(VERSION)/$(LINUX) $(LDFLAGS) ./cmd/$(BINARY)

darwin:
	env GOOS=darwin GOARCH=amd64 go build -v -o dist/$(VERSION)/$(DARWIN) $(LDFLAGS) ./cmd/$(BINARY)

windows:
	env GOOS=windows GOARCH=amd64 go build -v -o dist/$(VERSION)/$(WINDOWS) $(LDFLAGS) ./cmd/$(BINARY)

clean:
	rm -rf dist/$(VERSION)

.PHONY: all linux darwin windows