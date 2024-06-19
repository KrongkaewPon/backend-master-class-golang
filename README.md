# Backend Master Class [Golang + Postgres + Kubernetes + gRPC]

Ref: https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes

## Database Design

- Design a SQL DB schema using dbdiagram.io
- Save the DB schema as PDF/PNG diagram and share it with your team
- Generate SQL code to create the schema in a target database engine: Postgres

## Database Migration

Ref: https://github.com/golang-migrate/migrate

- [Installation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

## Generate CRUD Golang code from SQL

Ref: https://github.com/sqlc-dev/sqlc

- [Installation](https://docs.sqlc.dev/en/latest/overview/install.html)

## RESTful HTTP API

Ref: https://github.com/gin-gonic/gin

## Load Config

Ref: https://github.com/spf13/viper

## Mocking Framework

Ref: https://github.com/golang/mock

## [DBML - Database Markup Language](https://dbml.dbdiagram.io/home)

- vscode extension: vscode-dbml
- [cli](https://dbml.dbdiagram.io/cli)
  - [Convert a DBML file to SQL](https://dbml.dbdiagram.io/cli#convert-a-dbml-file-to-sql)
  - [Convert a SQL file to DBML](https://dbml.dbdiagram.io/cli#convert--a-sql-file-to-dbml)

## GRPC

- [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/)
- [Go plugins](https://grpc.io/docs/languages/go/quickstart/)

```sh
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

- [gRPC Client](https://github.com/ktr0731/evans)
- [gRPC-Gateway](https://github.com/grpc-ecosystem/grpc-gateway)
