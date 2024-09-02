package main

import (
	"github.com/TomasCruz/users/internal/app"
)

// @title Users
// @version 1.0
// @description Users service
// @contact.name TomasCruz
// @contact.url https://github.com/TomasCruz/users
// @license.name MIT
// @license.url https://mit-license.org/
func main() {
	// this approach of app.Start() is taken for two reasons:
	// 1) if e.g. a worker app for listening to Kafka is needed, add another directory and starter file like this one,
	// another directory like internal/app with appropriate content, proper Makefile commands, etc
	// 2) putting everything in internal
	app.Start()
}
