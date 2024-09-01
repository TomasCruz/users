package main

import (
	"os"
	"os/signal"

	"github.com/TomasCruz/users/internal/core"
	"github.com/TomasCruz/users/internal/handlers/httphandler"
	"github.com/TomasCruz/users/internal/repos/database"
	"github.com/TomasCruz/users/internal/repos/msg"
	"github.com/labstack/gommon/log"
)

func main() {
	// populate configuration
	config, err := setupFromEnvVars()
	if err != nil {
		log.Fatalf("failed to read environment variables: %s", err.Error())
	}

	// init DB
	db, err := database.InitDB(config)
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err.Error())
	}

	// Kafka producer
	msg, err := msg.InitMsg(config)
	if err != nil {
		log.Fatalf("failed to create Kafka producer: %s", err.Error())
	}

	// init HTTP handler
	e := httphandler.New(config, db, msg)

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	gracefulShutdown(db, msg, e)
}

func gracefulShutdown(db core.DB, msg core.Msg, e core.Http) {
	// DB
	defer db.Close()

	// Kafka
	defer msg.Close()

	// Echo
	e.Close()
}
