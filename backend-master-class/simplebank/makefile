DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

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
	migrate -path db/migration -database "$(DB_URL)" -verbose up

# Migrate up 1 step
migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

# Migrate down all steps
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

# Migrate down 1 step
migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

# Generate migration file
# e.g. `make new_migration name=create_users_table`
new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

# Run sqlc to generate code
sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

# Start server
server:
	go run main.go

# Generate mock code
mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/nuxzero/simplebank/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/nuxzero/simplebank/worker TaskDistributor

# Generate proto https://protobuf.dev/reference/go/go-generated/
proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
  proto/*.proto
	statik -src=./doc/swagger -dest=./doc

# Run evans cli for testing gRPC server
evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7.2.3-alpine

.PHONY: postgres createdb dropdb sqlc test server mock proto evans redis new_migration
