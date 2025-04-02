run:
	go run cmd/go-aws-cli/main.go

build:
	go build -o ./build/go-aws-cli cmd/go-aws-cli/main.go

.PHONY: run build