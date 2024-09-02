package main

import (
	"github.com/TomasCruz/users/internal/app"
)

// this approach of app.Start() is taken for two reasons:
// 1) if e.g. a worker app for listening to Kafka is needed, add another directory and starter file like this one,
// another directory like internal/app with appropriate content, proper Makefile commands, etc
// 2) putting everything in internal
func main() {
	app.Start()
}
