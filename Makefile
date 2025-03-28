GO=go

run:
	$(GO) run main.go

test:
	$(GO) test -v ./...

fmt:
	$(GO) fmt ./...

lint:
	golangci-lint run ./...

build:
	$(GO) build -o app

run-build:
	./app

clean:
	rm -f app
