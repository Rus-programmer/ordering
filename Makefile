include .env
export $(shell sed 's/=.*//' .env)

ENV_FILE := .env
ifeq (,$(wildcard $(ENV_FILE)))
$(error file .env not found)
endif

postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} -d postgres:17.2-alpine3.21

createdb:
	docker exec -it postgres17 createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}

dropdb:
	docker exec -it postgres17 dropdb --username=${DB_USER} ${DB_NAME}

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

migrateup:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down 1

.PHONY: postgres createdb dropdb new_migration migrateup migratedown migrateup1 migratedown1