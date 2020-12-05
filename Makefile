build:
	go build -o cmd/mlisp/mlisp cmd/mlisp/main.go

test:
	go test ./...
