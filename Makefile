build:
	go build -o cmd/tl/tl cmd/tl/main.go

test:
	go test ./...
