# Users
hexagonal architecture server example

## Features
- integrates [swaggo](https://github.com/swaggo/swag) and [echo-swagger](github.com/swaggo/echo-swagger) for generating swagger documentation and Swagger UI
- swagger documentation on http://localhost:8091/swagger/index.html
- initial Postgres database structure via [golang-migrate](https://github.com/golang-migrate/migrate)
- separate unit and integration testing
- separate docker compose and env files for running server and integration tests
- [mockery](https://github.com/vektra/mockery) for unit test mocks
- Makefile with various convenient targets
- provides gRPC and [NATS](github.com/nats-io/nats.go) inter-service communication
- supports message queues via [Kafka](https://github.com/confluentinc/confluent-kafka-go/)

## Running Prerequisites
- install Docker
- follow instructions on https://grpc.io/docs/languages/go/quickstart/ for gRPC
- install mockery from https://github.com/vektra/mockery/releases
- extract to /usr/local/go/bin/
- go install github.com/swaggo/swag/cmd/swag@latest
- if $HOME/go/bin is not in the PATH already, add it by running command 'export PATH=$PATH:$HOME/go/bin'

## Running instructions
- to run server, open the terminal and navigate to appropriate directory
- run 'make deps_up', then 'make', then 'make run'
- to run worker, open another terminal in the same directory
- run 'make wbuild', then 'make wrun'
- open another terminal in the same directory
- run (adjusted for IDs) curl commands from CURL.md for manually testing the endpoints
- run 'make integration' for integration testing
- run 'make test' for unit testing

## Project Structure

```
├── bin
├── cmd
|   ├── app
|   └── worker
├── docs
└── internal
|   ├─ core
|   |  ├── entities
|   |  ├── ports
|   |  └── service
|   |      ├── app
|   |      └─ worker
|   ├─ handlers
|   |  ├─ grpchandler
|   |  |  └─ users
|   |  └─ httphandler
|   └─ infra
|      ├── app
|      ├── configuration
|      ├── database
|      |   └─  migrations
|      ├── kafkaque
|      ├── log
|      ├── natsmsg
|      └── worker
└── tests
    ├── mocks
    └─ scripts
```

- **bin** - application binaries, gitignored
- **cmd** - entry points
  - **app** - REST/gRPC server entry point
  - **worker** - message queue worker entry point
- **docs** - swagger documentation generated files
- **internal** - code not exposed to external packages
  - **core** - business logic
    - **entities** - business concepts, inner core
    - **ports** -  service interfaces (inversion of control)
    - **service** - services for the applications
      - **app** - REST/gRPC server business logic
      - **worker** - message queue worker business logic
  - **handlers** - handlers (driver actors)
    - **grpchandler** - gRPC handlers
      - **users** - server's proto file and generated code (gitignored)
    - **httphandler** - REST handlers
  - **infra** - implementations (adapters) and infrastructure code
    - **app** - REST/gRPC server main file, instantiates and injects deps
    - **configuration** - configuration, shared between applications
    - **database** - DB implementation
      - **migrations** - DB migration scripts
    - **kafkaque** - Kafka message queue implementation
    - **log** - logger implementation (a cross-cutting concern)
    - **natsmsg** - NATS messaging implementation
    - **worker** - message queue worker main file
- **tests** - testing auxiliaries
  - **mocks** - generated mocks for unit testing
  - **scripts** - scripts used for integration testing

