include .env

AIR=~/go/bin/air
MIGRATE=~/go/bin/migrate


dev:
	$(AIR)
run:
	go run main.go
migrate-create:
	$(MIGRATE) create -ext sql -dir database/migrations $(name)
migrate-up:
	$(MIGRATE) -path database/migrations -database $(DATABASE_URL) up
migrate-down:
	$(MIGRATE) -path database/migrations -database $(DATABASE_URL) down