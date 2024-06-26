DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

startdb:
	docker compose up -d

stopdb:
	docker compose down

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen --build_flags=--mod=mod -destination db/mock/store.go simplebank/db/sqlc Store  

.PHONY: startdb stopdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock new_migration db_schema
