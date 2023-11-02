postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=12345 -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres12 dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:12345@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:12345@localhost:5432/simple_bank?sslmode=disable" -verbose down
migrateup1:
	migrate -path db/migration -database "postgresql://root:12345@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown1:
	migrate -path db/migration -database "postgresql://root:12345@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
test:
	go test ./... -v -cover
sqlc:
	sqlc generate
server:
	go run main.go
mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/rassulmagauin/simplebank/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown server mock