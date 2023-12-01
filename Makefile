.PHONY: run test build

run:
	@go run cmd/main.go

test:
	@go test -v ./pkg

build:
	@echo "unimplemented"