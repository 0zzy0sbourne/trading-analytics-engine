BIN := bin/tae

.PHONY: build run test vet fmt tidy clean

build:
	go build -o $(BIN) ./cmd/tae

run:
	go run ./cmd/tae

test:
	go test -race ./...

vet:
	go vet ./...

fmt:
	gofmt -s -w .

tidy:
	go mod tidy

clean:
	rm -rf bin
