all: clean build

.PHONY: clean
clean: fmt
	go clean
	rm -f bin/*

.PHONY: build
build:
	CGO_ENABLED=1 go build -o bin/server cmd/*

.PHONY: deps_up
deps_up:
	docker compose up -d

.PHONY: deps_down
deps_down:
	docker compose down -v

run:
	bin/server

fmt:
	gofmt -l -w -e ./

psql:
	docker run -it --rm --link pgdb:postgres --net users_default -e POSTGRES_USER=toma -e POSTGRES_PASSWORD=pswd -p 5351 postgres psql postgresql://toma:pswd@pgdb

qpsql:
	docker run -it --rm --link pgdb_test:postgres --net users_default_test -e POSTGRES_USER=test -e POSTGRES_PASSWORD=test -p 15351 postgres psql postgresql://test:test@pgdb_test

topics:
	docker compose exec kafka kafka-topics --create --topic user-created --partitions 1 --replication-factor 1 --bootstrap-server kafka:9092
	docker compose exec kafka kafka-topics --create --topic user-updated --partitions 1 --replication-factor 1 --bootstrap-server kafka:9092
	docker compose exec kafka kafka-topics --create --topic user-deleted --partitions 1 --replication-factor 1 --bootstrap-server kafka:9092

list-topics:
	docker compose exec kafka kafka-topics --bootstrap-server kafka:9092 --list

# Testing
.PHONY: mocks
mocks:
	rm -f ./tests/mocks/*.go
	CGO_ENABLED=1 /usr/local/go/bin/mockery --all --output ./tests/mocks --dir ./internal/

.PHONY: test
test: mocks docs fmt
	go test -v -count=1 -tags unit  ./...

.PHONY: intdeps_up
intdeps_up:
	docker compose -f docker-compose-test.yml up -d

.PHONY: inttests
inttests:
	go test -v -count=1 -tags integration ./...

.PHONY: intdeps_down
intdeps_down:
	docker compose -f docker-compose-test.yml down -v

.PHONY: integration
integration: intdeps_up inttests intdeps_down

# docs
.PHONY: docs
docs:
	swag init -g ./cmd/main.go -o ./docs/

# Docker stuff
img:
	docker build --tag dock-users .

bshimg:
	docker run -it dock-users bash

drun:
	docker run --net host dock-users
