deps:
	brew install ngrok
	go install github.com/cosmtrek/air@latest

up:
	docker-compose up -d
	air

down:
	docker-compose down

ps:
	docker-compose ps

migrate/up:
	go run migrations/main.go migrations/targets.go up

migrate/down:
	go run migrations/main.go migrations/targets.go down

seed/init:
	go run seeds/main.go seeds/targets.go init

seed/mock:
	go run seeds/main.go seeds/targets.go mock
