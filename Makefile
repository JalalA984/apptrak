APP_PATH = ./cmd/apptrak

.PHONY: fmt vet build exc clean test deps

default: exc

deps:
	go mod tidy
	go mod download

fmt: deps
	go fmt ./...

vet: fmt
	go vet ./...

test: vet
	go test -v ./internal/handlers/

build: test
	go build -o $(APP_PATH)/app $(APP_PATH)

exc: build
	$(APP_PATH)/app -port=":8080"

clean:
	go clean
	rm -f $(APP_PATH)/app
