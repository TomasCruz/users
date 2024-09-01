all: clean build

.PHONY: clean
clean: fmt
	go clean
	rm -f bin/*

.PHONY: build
build:
	CGO_ENABLED=1 go build -o bin/server cmd/*

run:
	bin/server

fmt:
	gofmt -l -w -e ./

psql:
	docker run -it --rm --link pgdb:postgres --net users_default -e POSTGRES_USER=toma -e POSTGRES_PASSWORD=pswd -p 5351 postgres psql postgresql://toma:pswd@pgdb

topics:
	docker compose exec kafka kafka-topics --create --topic user-created --partitions 1 --replication-factor 1 --bootstrap-server kafka:9092
	docker compose exec kafka kafka-topics --create --topic user-updated --partitions 1 --replication-factor 1 --bootstrap-server kafka:9092
	docker compose exec kafka kafka-topics --create --topic user-deleted --partitions 1 --replication-factor 1 --bootstrap-server kafka:9092

list-topics:
	docker compose exec kafka kafka-topics --bootstrap-server kafka:9092 --list
