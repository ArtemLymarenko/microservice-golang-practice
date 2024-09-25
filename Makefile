DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST=localhost
DB_PORT=5432

CREATE=migrate create -ext sql -dir $(path) -seq $(name)
create-migration:
	$(CREATE)

MIGRATE=migrate -source file://$(path) -database postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(db)?sslmode=disable
migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

