deps:
	go install github.com/cespare/reflex@latest

up:
	reflex -r '\.go|config.yml\z' -s -- sh -c 'go run cmd/main.go cmd/config.go cmd/router.go'
