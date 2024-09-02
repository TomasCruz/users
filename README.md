# Users

## Running Prerequisites
- install Docker
- install mockery from https://github.com/vektra/mockery/releases
- extract to /usr/local/go/bin/
- go install github.com/swaggo/swag/cmd/swag@latest
- if $HOME/go/bin is not in the PATH already, add it by running command 'export PATH=$PATH:$HOME/go/bin'

## Running instructions
- Open the terminal and navigate to appropriate directory
- run 'docker compose up -d'
- once it's ready, run 'make'
- run 'make run'
- open another terminal in the same directory
- run (adjusted for IDs and the like) curl commands from CURL.md for manually testing the endpoints
