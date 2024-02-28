.PHONY: database go-build go-run go-test
database:
	docker-compose up -d
build:
	go mod download && go build -o server main.go
test:
	go test ./...
run:
	./server