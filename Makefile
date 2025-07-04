postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple

dropdb:
	docker exec -it postgres17 dropdb --username=root  simple

migrateup:
	migrate -path db/migration -database "postgresql://root:1234@localhost:5432/simple?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:1234@localhost:5432/simple?sslmode=disable" -verbose  down

sqlc:
	sqlc generate

server:
	go run main.go

test:
	go test -v -cover ./...
.PHONY: createdb dropdb postgres migrateup migratedown sqlc server
