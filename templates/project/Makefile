up_docker:
	docker-compose up -d

down:
	docker-compose down

up: up_docker
	go run main.go start --env-file=.env.dev

test:
	go test -race -v ./...