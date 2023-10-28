postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createDb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropDb:
	docker exec -it postgres12 dropdb simple_bank

migrateUp:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateDown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server-run:
	go run ./cmd/web/

server-build:
	go build ./cmd/web/

clean:
	go mod tidy

test:
	go test -v -cover ./...

mock:
	mockgen -package mockdb -destignation db/mock/store.go github.com/kararnab/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock