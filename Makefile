run:
	go run main.go

build:
	go build -o ./build/app main.go

.PHONY: run build