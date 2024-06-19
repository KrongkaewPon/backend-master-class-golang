# Simple Bank

## Database Migration

Create a migration.

```sh
migrate create -ext sql -dir db/migration -seq init_schema
```

Run migrations.

```sh
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
```

## Generate CRUD code from SQL

Write a sql query in `db/query`, then run `make sqlc`.

## Production

- Set `GIN_MODE=release`
