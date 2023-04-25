include .env
export

.PHONY: run-app
run-app:
	go run ./cmd/app/main.go

.PHONY: tidy
tidy:
	go mod tidy