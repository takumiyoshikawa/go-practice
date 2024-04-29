postgres-init:
	docker run --rm -d --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password postgres:latest

postgres:
	docker exec -it postgres psql

createdb:
	docker exec -it postgres createdb --username=root --owner=root go-chat

dropdb:
	docker exec -it postgres dropdb go-chat

migrate-up:
	

.PHONY: postgres-init