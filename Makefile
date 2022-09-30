postgres:
	docker run --name postgres14 -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -p 5431:5432 -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank

createtestdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank_test

startdb:
	docker start postgres14

dropdb:
	docker exec -it postgres14 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5431/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5431/simple_bank?sslmode=disable" -verbose down

migratetestup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5431/simple_bank_test?sslmode=disable" -verbose up

migratetestdown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5431/simple_bank_test?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb startdb dropdb migrateup migratedown sqlc test
