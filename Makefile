run:
	go run cmd/aws-cli/main.go

build:
	go build -o ./build/app cmd/aws-cli/main.go

.PHONY: run build