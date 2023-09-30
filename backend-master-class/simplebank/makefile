# Create Postgres container
postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

# Create new database
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

# Drop database
dropdb:
	docker exec -it postgres12 dropdb simple_bank	

# Migrate up all steps
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

# Migrate up 1 step
migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

# Migrate down all steps
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

# Migrate down 1 step
migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

# Run sqlc to generate code
sqlc:
	sqlc generate

test:
	go test -v -cover ./...

# Start server
server:
	go run main.go

# Generate mock code
mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/nuxzero/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb sqlc test server mock
