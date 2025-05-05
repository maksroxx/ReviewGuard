APP_NAME=reviewguard

.PHONY: build run test docker-up docker-down clean

build:
	go build -o bin/$(APP_NAME) ./cmd/

run: build
	./bin/$(APP_NAME)

test:
	go test ./...

docker-up:
	docker-compose up --build -d

docker-down:
	docker-compose down

clean:
	rm -f bin/$(APP_NAME)
