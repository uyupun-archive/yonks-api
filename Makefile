deps:
	go install github.com/cosmtrek/air@latest

up:
	docker-compose up -d
	air

down:
	docker-compose down

ps:
	docker-compose ps

batch:
	go run batches/scheduler.go batches/matcher.go

migrate/up:
	go run migrations/migrator.go migrations/targets.go up

migrate/down:
	go run migrations/migrator.go migrations/targets.go down

seed/init:
	go run seeds/seeder.go seeds/targets.go init

seed/mock:
	go run seeds/seeder.go seeds/targets.go mock
