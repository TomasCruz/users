package app

import (
	"os"
	"os/signal"

	"github.com/TomasCruz/users/internal/core"
	"github.com/TomasCruz/users/internal/database"
	"github.com/TomasCruz/users/internal/handlers/httphandler"
	"github.com/TomasCruz/users/internal/msg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Start() {
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

	// new Core
	cr := core.New(config, db, msg)

	// init HTTP handler
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	h := httphandler.New(e, cr, config)

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	gracefulShutdown(db, msg, h)
}

func gracefulShutdown(db core.DB, msg core.Msg, h httphandler.HTTPHandler) {
	// Echo
	err := h.Close()
	if err != nil {
		log.Error(err)
	}

	// Kafka
	msg.Close()

	// DB
	err = db.Close()
	if err != nil {
		log.Error(err)
	}
}
