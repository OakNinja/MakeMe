BINARY_NAME=mm

.PHONY: all build test clean install

all: build

build:
	go build -o $(BINARY_NAME) ./cmd/makemego

test:
	go test ./...

clean:
	go clean
	rm -f $(BINARY_NAME)

install:
	go build -o $(BINARY_NAME) ./cmd/makemego
	sudo mv $(BINARY_NAME) /usr/local/bin/
