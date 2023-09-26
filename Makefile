# Use for lite start

run:
	@go run ./cmd/app

build:
	@go build ./cmd/app

test:
	@go test -v ./internal/pkg/util/