up_docker:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs

watch-logs:
	docker-compose logs -f

up_app:
	go run main.go start --env-file=.env.dev

up: up_docker up_app

test:
	go test -race -v ./...
