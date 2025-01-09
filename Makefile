main:
	docker compose run app go run main.go

setup:
	docker compose run app go run cmd/setup/main.go

seed:
	docker compose run app go run cmd/seed/main.go
