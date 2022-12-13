BIN := bin/cs

.PHONY: all
all: build

.PHONY: build
build: test clean
	go build -o $(BIN) -v ./cmd/cinema-search

.PHONY: clean
clean:
	rm -rf $(BIN)
	go clean

.PHONY: test
test:
	go test ./... -count=1 --race -v
