DB_USER ?= $(shell echo $$DB_USER)
DB_PASSWORD ?= $(shell echo $$DB_PASSWORD)
DB_HOST ?= $(shell echo $$DB_HOST)
DB_PORT ?= $(shell echo $$DB_PORT)
MIGRATE=migrate -source file://$(path) -database postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(table)?sslmode=disable
CREATE=migrate create -ext sql -dir $(path) -seq $(name)

create-migration:
	$(CREATE)

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

