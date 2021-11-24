.PHONY: deps up ngrok

deps:
	brew install ngrok
	go install github.com/cespare/reflex@latest

up:
	reflex -r '\.go|config.yml\z' -s -- sh -c 'go run cmd/main.go cmd/config.go cmd/router.go'

ngrok:
	ngrok http 8080

migrate/up:
	go run migrations/main.go migrations/targets.go up

migrate/down:
	go run migrations/main.go migrations/targets.go down

seed/init:
	go run seeds/main.go seeds/targets.go init

seed/mock:
	go run seeds/main.go seeds/targets.go mock
