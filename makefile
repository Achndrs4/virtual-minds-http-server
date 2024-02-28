default: fmt build-server run-server

start-db:
	docker-compose -f database.yml up -d --build --remove-orphans
kill-db:
	docker-compose -f database.yml down --rmi "local"
build-server:
	go build . 
build-docker:
	docker build . -t virtual-minds/http-server --no-cache
run-docker-server:
	docker run --name server virtual-minds/http-server 
kill-docker-server:
	docker rm server
fmt:
	go mod tidy && go fmt ./...
test:
	go test ./...
run-server:
	./http-server
