postgres:
	docker run --name pdb -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=toor -d postgres:12-alpine

createdb:
	docker exec -it pdb createdb --username=root --owner=root simple_bank

setup:
	make postgres
	make createdb
	make migrate-up
	make clean

dropdb:
	docker exec -it pdb dropdb simple_bank

stop:
	docker stop pdb

start:
	docker start pdb

clean:
	docker stop pdb
	docker rm pdb

migrate-up:
	migrate -path db/migration -database "postgresql://root:toor@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://root:toor@localhost:5432/simple_bank?sslmode=disable" -verbose down -all

sqlc: 
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb stop-server start-server clean check-image migrate-up migrate-down sqlc go-testmain

