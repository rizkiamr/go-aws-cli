run:
	go run cmd/aws-cli/main.go

build:
	go build -o ./build/aws-cli cmd/aws-cli/main.go

.PHONY: run build