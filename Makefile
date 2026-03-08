BINARY := shell-imagen

.PHONY: build test vet clean

build:
	go build -o $(BINARY) ./cmd/shell-imagen

test:
	go test ./...

vet:
	go vet ./...

clean:
	rm -f $(BINARY)
