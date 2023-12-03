.PHONY: run test build

run:
	@go run cmd/main.go

test:
	@go test -v ./pkg

benchmark:
	@go test -bench=. ./pkg

build:
	@echo "unimplemented"