# Simple bank app

This is the practice project for the master class backend course.
[link](https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/learn/lecture/31063566?start=0#overview)

## Setup

- Start Postgres container `make postgres`
- If the postgres container already created you can just start the container `docker start postgres12`
- Create new database `make createdb`
- Run migrations `make migrateup`
- Start the server `make server`

## SQLC

In this project we use [sqlc](https://github.com/sqlc-dev/sqlc) to generate Go code from SQL queries.
When you change the queries in `db/query` directory, you need to run `make sqlc` to regenerate the code.
The generated code is in `db/sqlc` directory and the file ends with `*.sql.go`.

## Run unit tests

```sh
make test
```

## Stop Postgres container

```sh
docker stop postgres12
```

## Database Migration

### Create new migration

```sh
migrate create -ext sql -dir db/migration -seq <migration_name>
# Example
migrate create -ext sql -dir db/migration -seq add_sessions
```

### Fix and force version error

In case fail to migration up with error `error: Dirty database version 3. Fix and force version.`.
You can force the version by running the command below. Then run migrate up again. [issue link](https://github.com/golang-migrate/migrate/issues/282)

```sh
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" force 2 
```

## gRPC Testing

Use [evans](https://github.com/ktr0731/evans) to test gRPC from command line.

### Usage

Run evans

```sh
evans --host localhost --port 9090 -r repl
```

Show all services

```sh
show service
```

Call RPC method

```sh
call CreateUser
```

## Documentation

You can open http://localhost:8080/swagger/ to check out API documentation.
